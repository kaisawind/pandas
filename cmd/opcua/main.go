// Copyright (c) Mainflux
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/cloustone/pandas"
	"github.com/cloustone/pandas/mainflux/broker"
	"github.com/cloustone/pandas/mainflux/opcua"
	"github.com/cloustone/pandas/mainflux/opcua/api"
	"github.com/cloustone/pandas/mainflux/opcua/db"
	"github.com/cloustone/pandas/mainflux/opcua/gopcua"
	"github.com/cloustone/pandas/mainflux/opcua/redis"
	"github.com/cloustone/pandas/pkg/logger"
	r "github.com/go-redis/redis"

	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

const (
	defHTTPPort       = "8188"
	defOPCIntervalMs  = "1000"
	defOPCPolicy      = ""
	defOPCMode        = ""
	defOPCCertFile    = ""
	defOPCKeyFile     = ""
	defNatsURL        = pandas.DefNatsURL
	defLogLevel       = "debug"
	defESURL          = "localhost:6379"
	defESPass         = ""
	defESDB           = "0"
	defESConsumerName = "opcua"
	defRouteMapURL    = "localhost:6379"
	defRouteMapPass   = ""
	defRouteMapDB     = "0"

	envHTTPPort       = "PD_OPCUA_ADAPTER_HTTP_PORT"
	envLogLevel       = "PD_OPCUA_ADAPTER_LOG_LEVEL"
	envOPCIntervalMs  = "PD_OPCUA_ADAPTER_INTERVAL_MS"
	envOPCPolicy      = "PD_OPCUA_ADAPTER_POLICY"
	envOPCMode        = "PD_OPCUA_ADAPTER_MODE"
	envOPCCertFile    = "PD_OPCUA_ADAPTER_CERT_FILE"
	envOPCKeyFile     = "PD_OPCUA_ADAPTER_KEY_FILE"
	envNatsURL        = "PD_NATS_URL"
	envESURL          = "PD_THINGS_ES_URL"
	envESPass         = "PD_THINGS_ES_PASS"
	envESDB           = "PD_THINGS_ES_DB"
	envESConsumerName = "PD_OPCUA_ADAPTER_EVENT_CONSUMER"
	envRouteMapURL    = "PD_OPCUA_ADAPTER_ROUTE_MAP_URL"
	envRouteMapPass   = "PD_OPCUA_ADAPTER_ROUTE_MAP_PASS"
	envRouteMapDB     = "PD_OPCUA_ADAPTER_ROUTE_MAP_DB"

	thingsRMPrefix     = "thing"
	channelsRMPrefix   = "channel"
	connectionRMPrefix = "connection"
)

type config struct {
	httpPort       string
	opcuaConfig    opcua.Config
	natsURL        string
	logLevel       string
	esURL          string
	esPass         string
	esDB           string
	esConsumerName string
	routeMapURL    string
	routeMapPass   string
	routeMapDB     string
}

func main() {
	cfg := loadConfig()

	logger, err := logger.New(os.Stdout, cfg.logLevel)
	if err != nil {
		log.Fatalf(err.Error())
	}

	rmConn := connectToRedis(cfg.routeMapURL, cfg.routeMapPass, cfg.routeMapDB, logger)
	defer rmConn.Close()

	thingRM := newRouteMapRepositoy(rmConn, thingsRMPrefix, logger)
	chanRM := newRouteMapRepositoy(rmConn, channelsRMPrefix, logger)
	connRM := newRouteMapRepositoy(rmConn, connectionRMPrefix, logger)

	esConn := connectToRedis(cfg.esURL, cfg.esPass, cfg.esDB, logger)
	defer esConn.Close()

	b, err := broker.New(cfg.natsURL)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer b.Close()

	ctx := context.Background()
	sub := gopcua.NewSubscriber(ctx, b, thingRM, chanRM, connRM, logger)
	browser := gopcua.NewBrowser(ctx, logger)

	svc := opcua.New(sub, browser, thingRM, chanRM, connRM, cfg.opcuaConfig, logger)
	svc = api.LoggingMiddleware(svc, logger)
	svc = api.MetricsMiddleware(
		svc,
		kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "opc_adapter",
			Subsystem: "api",
			Name:      "request_count",
			Help:      "Number of requests received.",
		}, []string{"method"}),
		kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "opc_adapter",
			Subsystem: "api",
			Name:      "request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, []string{"method"}),
	)

	go subscribeToStoredSubs(sub, cfg.opcuaConfig, logger)
	go subscribeToThingsES(svc, esConn, cfg.esConsumerName, logger)

	errs := make(chan error, 2)

	go startHTTPServer(svc, cfg, logger, errs)

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	err = <-errs
	logger.Error(fmt.Sprintf("OPC-UA adapter terminated: %s", err))
}

func loadConfig() config {
	oc := opcua.Config{
		Interval: pandas.Env(envOPCIntervalMs, defOPCIntervalMs),
		Policy:   pandas.Env(envOPCPolicy, defOPCPolicy),
		Mode:     pandas.Env(envOPCMode, defOPCMode),
		CertFile: pandas.Env(envOPCCertFile, defOPCCertFile),
		KeyFile:  pandas.Env(envOPCKeyFile, defOPCKeyFile),
	}
	return config{
		httpPort:       pandas.Env(envHTTPPort, defHTTPPort),
		opcuaConfig:    oc,
		natsURL:        pandas.Env(envNatsURL, defNatsURL),
		logLevel:       pandas.Env(envLogLevel, defLogLevel),
		esURL:          pandas.Env(envESURL, defESURL),
		esPass:         pandas.Env(envESPass, defESPass),
		esDB:           pandas.Env(envESDB, defESDB),
		esConsumerName: pandas.Env(envESConsumerName, defESConsumerName),
		routeMapURL:    pandas.Env(envRouteMapURL, defRouteMapURL),
		routeMapPass:   pandas.Env(envRouteMapPass, defRouteMapPass),
		routeMapDB:     pandas.Env(envRouteMapDB, defRouteMapDB),
	}
}

func connectToRedis(redisURL, redisPass, redisDB string, logger logger.Logger) *r.Client {
	db, err := strconv.Atoi(redisDB)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to connect to redis: %s", err))
		os.Exit(1)
	}

	return r.NewClient(&r.Options{
		Addr:     redisURL,
		Password: redisPass,
		DB:       db,
	})
}

func subscribeToStoredSubs(sub opcua.Subscriber, cfg opcua.Config, logger logger.Logger) {
	// Get all stored subscriptions
	nodes, err := db.ReadAll()
	if err != nil {
		logger.Warn(fmt.Sprintf("Read stored subscriptions failed: %s", err))
	}

	for _, n := range nodes {
		cfg.ServerURI = n.ServerURI
		cfg.NodeID = n.NodeID
		go func() {
			if err := sub.Subscribe(cfg); err != nil {
				logger.Warn(fmt.Sprintf("Subscription failed: %s", err))
			}
		}()
	}
}

func subscribeToOpcuaServer(gc opcua.Subscriber, cfg opcua.Config, logger logger.Logger) {
	if err := gc.Subscribe(cfg); err != nil {
		logger.Warn(fmt.Sprintf("OPC-UA Subscription failed: %s", err))
	}
}

func subscribeToThingsES(svc opcua.Service, client *r.Client, prefix string, logger logger.Logger) {
	eventStore := redis.NewEventStore(svc, client, prefix, logger)
	if err := eventStore.Subscribe("mainflux.things"); err != nil {
		logger.Warn(fmt.Sprintf("Failed to subscribe to Redis event source: %s", err))
	}
}

func newRouteMapRepositoy(client *r.Client, prefix string, logger logger.Logger) opcua.RouteMapRepository {
	logger.Info(fmt.Sprintf("Connected to %s Redis Route-map", prefix))
	return redis.NewRouteMapRepository(client, prefix)
}

func startHTTPServer(svc opcua.Service, cfg config, logger logger.Logger, errs chan error) {
	p := fmt.Sprintf(":%s", cfg.httpPort)
	logger.Info(fmt.Sprintf("opcua-adapter service started, exposed port %s", cfg.httpPort))
	errs <- http.ListenAndServe(p, api.MakeHandler(svc))
}
