// Copyright (c) Mainflux
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/cloustone/pandas"
	"github.com/cloustone/pandas/mainflux/broker"
	"github.com/cloustone/pandas/mainflux/transformers/senml"
	"github.com/cloustone/pandas/mainflux/writers"
	"github.com/cloustone/pandas/mainflux/writers/api"
	"github.com/cloustone/pandas/mainflux/writers/influxdb"
	"github.com/cloustone/pandas/pkg/logger"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	influxdata "github.com/influxdata/influxdb/client/v2"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

const (
	svcName = "influxdb-writer"

	defNatsURL         = pandas.DefNatsURL
	defLogLevel        = "error"
	defPort            = "8180"
	defDBName          = "mainflux"
	defDBHost          = "localhost"
	defDBPort          = "8086"
	defDBUser          = "mainflux"
	defDBPass          = "mainflux"
	defSubjectsCfgPath = "/config/subjects.toml"

	envNatsURL         = "PD_NATS_URL"
	envLogLevel        = "PD_INFLUX_WRITER_LOG_LEVEL"
	envPort            = "PD_INFLUX_WRITER_PORT"
	envDBName          = "PD_INFLUX_WRITER_DB_NAME"
	envDBHost          = "PD_INFLUX_WRITER_DB_HOST"
	envDBPort          = "PD_INFLUX_WRITER_DB_PORT"
	envDBUser          = "PD_INFLUX_WRITER_DB_USER"
	envDBPass          = "PD_INFLUX_WRITER_DB_PASS"
	envSubjectsCfgPath = "PD_INFLUX_WRITER_SUBJECTS_CONFIG"
)

type config struct {
	natsURL         string
	logLevel        string
	port            string
	dbName          string
	dbHost          string
	dbPort          string
	dbUser          string
	dbPass          string
	subjectsCfgPath string
}

func main() {
	cfg, clientCfg := loadConfigs()

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

	client, err := influxdata.NewHTTPClient(clientCfg)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to create InfluxDB client: %s", err))
		os.Exit(1)
	}
	defer client.Close()

	repo := influxdb.New(client, cfg.dbName)

	counter, latency := makeMetrics()
	repo = api.LoggingMiddleware(repo, logger)
	repo = api.MetricsMiddleware(repo, counter, latency)
	st := senml.New()
	if err := writers.Start(b, repo, st, svcName, cfg.subjectsCfgPath, logger); err != nil {
		logger.Error(fmt.Sprintf("Failed to start InfluxDB writer: %s", err))
		os.Exit(1)
	}

	errs := make(chan error, 2)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go startHTTPService(cfg.port, logger, errs)

	err = <-errs
	logger.Error(fmt.Sprintf("InfluxDB writer service terminated: %s", err))
}

func loadConfigs() (config, influxdata.HTTPConfig) {
	cfg := config{
		natsURL:         pandas.Env(envNatsURL, defNatsURL),
		logLevel:        pandas.Env(envLogLevel, defLogLevel),
		port:            pandas.Env(envPort, defPort),
		dbName:          pandas.Env(envDBName, defDBName),
		dbHost:          pandas.Env(envDBHost, defDBHost),
		dbPort:          pandas.Env(envDBPort, defDBPort),
		dbUser:          pandas.Env(envDBUser, defDBUser),
		dbPass:          pandas.Env(envDBPass, defDBPass),
		subjectsCfgPath: pandas.Env(envSubjectsCfgPath, defSubjectsCfgPath),
	}

	clientCfg := influxdata.HTTPConfig{
		Addr:     fmt.Sprintf("http://%s:%s", cfg.dbHost, cfg.dbPort),
		Username: cfg.dbUser,
		Password: cfg.dbPass,
	}

	return cfg, clientCfg
}

func makeMetrics() (*kitprometheus.Counter, *kitprometheus.Summary) {
	counter := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "influxdb",
		Subsystem: "message_writer",
		Name:      "request_count",
		Help:      "Number of database inserts.",
	}, []string{"method"})

	latency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "influxdb",
		Subsystem: "message_writer",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of inserts in microseconds.",
	}, []string{"method"})

	return counter, latency
}

func startHTTPService(port string, logger logger.Logger, errs chan error) {
	p := fmt.Sprintf(":%s", port)
	logger.Info(fmt.Sprintf("InfluxDB writer service started, exposed port %s", p))
	errs <- http.ListenAndServe(p, api.MakeHandler(svcName))
}
