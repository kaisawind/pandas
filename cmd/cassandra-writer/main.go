// Copyright (c) Mainflux
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	"github.com/gocql/gocql"
	"github.com/cloustone/pandas/mainflux"
	"github.com/cloustone/pandas/mainflux/broker"
	"github.com/cloustone/pandas/mainflux/logger"
	"github.com/cloustone/pandas/mainflux/transformers/senml"
	"github.com/cloustone/pandas/mainflux/writers"
	"github.com/cloustone/pandas/mainflux/writers/api"
	"github.com/cloustone/pandas/mainflux/writers/cassandra"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

const (
	svcName = "cassandra-writer"
	sep     = ","

	defNatsURL         = mainflux.DefNatsURL
	defLogLevel        = "error"
	defPort            = "8180"
	defCluster         = "127.0.0.1"
	defKeyspace        = "mainflux"
	defDBUsername      = ""
	defDBPassword      = ""
	defDBPort          = "9042"
	defSubjectsCfgPath = "/config/subjects.toml"

	envNatsURL         = "MF_NATS_URL"
	envLogLevel        = "MF_CASSANDRA_WRITER_LOG_LEVEL"
	envPort            = "MF_CASSANDRA_WRITER_PORT"
	envCluster         = "MF_CASSANDRA_WRITER_DB_CLUSTER"
	envKeyspace        = "MF_CASSANDRA_WRITER_DB_KEYSPACE"
	envDBUsername      = "MF_CASSANDRA_WRITER_DB_USERNAME"
	envDBPassword      = "MF_CASSANDRA_WRITER_DB_PASSWORD"
	envDBPort          = "MF_CASSANDRA_WRITER_DB_PORT"
	envSubjectsCfgPath = "MF_CASSANDRA_WRITER_SUBJECTS_CONFIG"
)

type config struct {
	natsURL         string
	logLevel        string
	port            string
	dbCfg           cassandra.DBConfig
	subjectsCfgPath string
}

func main() {
	cfg := loadConfig()

	logger, err := logger.New(os.Stdout, cfg.logLevel)
	if err != nil {
		log.Fatalf(err.Error())
	}

	b, err := broker.New(cfg.natsURL)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer b.Close()

	session := connectToCassandra(cfg.dbCfg, logger)
	defer session.Close()

	repo := newService(session, logger)
	st := senml.New()
	if err := writers.Start(b, repo, st, svcName, cfg.subjectsCfgPath, logger); err != nil {
		logger.Error(fmt.Sprintf("Failed to create Cassandra writer: %s", err))
	}

	errs := make(chan error, 2)

	go startHTTPServer(cfg.port, errs, logger)

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	err = <-errs
	logger.Error(fmt.Sprintf("Cassandra writer service terminated: %s", err))
}

func loadConfig() config {
	dbPort, err := strconv.Atoi(mainflux.Env(envDBPort, defDBPort))
	if err != nil {
		log.Fatal(err)
	}

	dbCfg := cassandra.DBConfig{
		Hosts:    strings.Split(mainflux.Env(envCluster, defCluster), sep),
		Keyspace: mainflux.Env(envKeyspace, defKeyspace),
		Username: mainflux.Env(envDBUsername, defDBUsername),
		Password: mainflux.Env(envDBPassword, defDBPassword),
		Port:     dbPort,
	}

	return config{
		natsURL:         mainflux.Env(envNatsURL, defNatsURL),
		logLevel:        mainflux.Env(envLogLevel, defLogLevel),
		port:            mainflux.Env(envPort, defPort),
		dbCfg:           dbCfg,
		subjectsCfgPath: mainflux.Env(envSubjectsCfgPath, defSubjectsCfgPath),
	}
}

func connectToCassandra(dbCfg cassandra.DBConfig, logger logger.Logger) *gocql.Session {
	session, err := cassandra.Connect(dbCfg)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to connect to Cassandra cluster: %s", err))
		os.Exit(1)
	}

	return session
}

func newService(session *gocql.Session, logger logger.Logger) writers.MessageRepository {
	repo := cassandra.New(session)
	repo = api.LoggingMiddleware(repo, logger)
	repo = api.MetricsMiddleware(
		repo,
		kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "cassandra",
			Subsystem: "message_writer",
			Name:      "request_count",
			Help:      "Number of requests received.",
		}, []string{"method"}),
		kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "cassandra",
			Subsystem: "message_writer",
			Name:      "request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, []string{"method"}),
	)

	return repo
}

func startHTTPServer(port string, errs chan error, logger logger.Logger) {
	p := fmt.Sprintf(":%s", port)
	logger.Info(fmt.Sprintf("Cassandra writer service started, exposed port %s", port))
	errs <- http.ListenAndServe(p, api.MakeHandler(svcName))
}