package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/cloustone/pandas"
	"github.com/cloustone/pandas/mainflux/broker"
	mqtt "github.com/cloustone/pandas/mainflux/mqtt"
	mr "github.com/cloustone/pandas/mainflux/mqtt/redis"
	"github.com/cloustone/pandas/pkg/logger"
	thingsapi "github.com/cloustone/pandas/things/api/auth/grpc"
	"github.com/go-redis/redis"
	mp "github.com/mainflux/mproxy/pkg/mqtt"
	ws "github.com/mainflux/mproxy/pkg/websocket"
	opentracing "github.com/opentracing/opentracing-go"
	jconfig "github.com/uber/jaeger-client-go/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	// MQTT
	defMQTTHost       = "0.0.0.0"
	defMQTTPort       = "1883"
	defMQTTTargetHost = "0.0.0.0"
	defMQTTTargetPort = "1883"
	envMQTTHost       = "PD_MQTT_ADAPTER_MQTT_HOST"
	envMQTTPort       = "PD_MQTT_ADAPTER_MQTT_PORT"
	envMQTTTargetHost = "PD_MQTT_ADAPTER_MQTT_TARGET_HOST"
	envMQTTTargetPort = "PD_MQTT_ADAPTER_MQTT_TARGET_PORT"
	// HTTP
	defHTTPHost       = "0.0.0.0"
	defHTTPPort       = "8080"
	defHTTPScheme     = "ws"
	defHTTPTargetHost = "localhost"
	defHTTPTargetPort = "8080"
	defHTTPTargetPath = "/mqtt"
	envHTTPHost       = "PD_MQTT_ADAPTER_WS_HOST"
	envHTTPPort       = "PD_MQTT_ADAPTER_WS_PORT"
	envHTTPScheme     = "PD_MQTT_ADAPTER_WS_SCHEMA"
	envHTTPTargetHost = "PD_MQTT_ADAPTER_WS_TARGET_HOST"
	envHTTPTargetPort = "PD_MQTT_ADAPTER_WS_TARGET_PORT"
	envHTTPTargetPath = "PD_MQTT_ADAPTER_WS_TARGET_PATH"
	// Logging
	defLogLevel = "error"
	envLogLevel = "PD_MQTT_ADAPTER_LOG_LEVEL"
	// Things
	defThingsURL     = "localhost:8181"
	defThingsTimeout = "1" // in seconds
	envThingsURL     = "PD_THINGS_URL"
	envThingsTimeout = "PD_MQTT_ADAPTER_THINGS_TIMEOUT"
	// Nats
	defNatsURL = pandas.DefNatsURL
	envNatsURL = "PD_NATS_URL"
	// Jaeger
	defJaegerURL = ""
	envJaegerURL = "PD_JAEGER_URL"
	// TLS
	defClientTLS = "false"
	defCACerts   = ""
	envClientTLS = "PD_MQTT_ADAPTER_CLIENT_TLS"
	envCACerts   = "PD_MQTT_ADAPTER_CA_CERTS"
	// Instance
	envInstance = "PD_MQTT_ADAPTER_INSTANCE"
	defInstance = ""
	// ES
	envESURL  = "PD_MQTT_ADAPTER_ES_URL"
	envESPass = "PD_MQTT_ADAPTER_ES_PASS"
	envESDB   = "PD_MQTT_ADAPTER_ES_DB"
	defESURL  = "localhost:6379"
	defESPass = ""
	defESDB   = "0"
)

type config struct {
	mqttHost       string
	mqttPort       string
	mqttTargetHost string
	mqttTargetPort string
	httpHost       string
	httpPort       string
	httpScheme     string
	httpTargetHost string
	httpTargetPort string
	httpTargetPath string
	jaegerURL      string
	logLevel       string
	thingsURL      string
	thingsTimeout  time.Duration
	natsURL        string
	clientTLS      bool
	caCerts        string
	instance       string
	esURL          string
	esPass         string
	esDB           string
}

func main() {
	cfg := loadConfig()

	logger, err := logger.New(os.Stdout, cfg.logLevel)
	if err != nil {
		log.Fatalf(err.Error())
	}

	conn := connectToThings(cfg, logger)
	defer conn.Close()

	tracer, closer := initJaeger("mproxy", cfg.jaegerURL, logger)
	defer closer.Close()

	thingsTracer, thingsCloser := initJaeger("things", cfg.jaegerURL, logger)
	defer thingsCloser.Close()

	rc := connectToRedis(cfg.esURL, cfg.esPass, cfg.esDB, logger)
	defer rc.Close()

	cc := thingsapi.NewClient(conn, thingsTracer, cfg.thingsTimeout)

	b, err := broker.New(cfg.natsURL)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer b.Close()

	es := mr.NewEventStore(rc, cfg.instance)

	// Event handler for MQTT hooks
	evt := mqtt.New(b, cc, es, logger, tracer)

	errs := make(chan error, 2)

	logger.Info(fmt.Sprintf("Starting MQTT proxy on port %s", cfg.mqttPort))
	go proxyMQTT(cfg, logger, evt, errs)

	logger.Info(fmt.Sprintf("Starting MQTT over WS  proxy on port %s", cfg.httpPort))
	go proxyWS(cfg, logger, evt, errs)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	err = <-errs
	logger.Error(fmt.Sprintf("mProxy terminated: %s", err))
}

func loadConfig() config {
	tls, err := strconv.ParseBool(pandas.Env(envClientTLS, defClientTLS))
	if err != nil {
		log.Fatalf("Invalid value passed for %s\n", envClientTLS)
	}

	timeout, err := strconv.ParseInt(pandas.Env(envThingsTimeout, defThingsTimeout), 10, 64)
	if err != nil {
		log.Fatalf("Invalid %s value: %s", envThingsTimeout, err.Error())
	}

	return config{
		mqttHost:       pandas.Env(envMQTTHost, defMQTTHost),
		mqttPort:       pandas.Env(envMQTTPort, defMQTTPort),
		mqttTargetHost: pandas.Env(envMQTTTargetHost, defMQTTTargetHost),
		mqttTargetPort: pandas.Env(envMQTTTargetPort, defMQTTTargetPort),
		httpHost:       pandas.Env(envHTTPHost, defHTTPHost),
		httpPort:       pandas.Env(envHTTPPort, defHTTPPort),
		httpScheme:     pandas.Env(envHTTPScheme, defHTTPScheme),
		httpTargetHost: pandas.Env(envHTTPTargetHost, defHTTPTargetHost),
		httpTargetPort: pandas.Env(envHTTPTargetPort, defHTTPTargetPort),
		httpTargetPath: pandas.Env(envHTTPTargetPath, defHTTPTargetPath),
		jaegerURL:      pandas.Env(envJaegerURL, defJaegerURL),
		thingsTimeout:  time.Duration(timeout) * time.Second,
		thingsURL:      pandas.Env(envThingsURL, defThingsURL),
		natsURL:        pandas.Env(envNatsURL, defNatsURL),
		logLevel:       pandas.Env(envLogLevel, defLogLevel),
		clientTLS:      tls,
		caCerts:        pandas.Env(envCACerts, defCACerts),
		instance:       pandas.Env(envInstance, defInstance),
		esURL:          pandas.Env(envESURL, defESURL),
		esPass:         pandas.Env(envESPass, defESPass),
		esDB:           pandas.Env(envESDB, defESDB),
	}
}

func initJaeger(svcName, url string, logger logger.Logger) (opentracing.Tracer, io.Closer) {
	if url == "" {
		return opentracing.NoopTracer{}, ioutil.NopCloser(nil)
	}

	tracer, closer, err := jconfig.Configuration{
		ServiceName: svcName,
		Sampler: &jconfig.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &jconfig.ReporterConfig{
			LocalAgentHostPort: url,
			LogSpans:           true,
		},
	}.NewTracer()
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to init Jaeger client: %s", err))
		os.Exit(1)
	}

	return tracer, closer
}

func connectToThings(cfg config, logger logger.Logger) *grpc.ClientConn {
	var opts []grpc.DialOption
	if cfg.clientTLS {
		if cfg.caCerts != "" {
			tpc, err := credentials.NewClientTLSFromFile(cfg.caCerts, "")
			if err != nil {
				logger.Error(fmt.Sprintf("Failed to load certs: %s", err))
				os.Exit(1)
			}
			opts = append(opts, grpc.WithTransportCredentials(tpc))
		}
	} else {
		logger.Info("gRPC communication is not encrypted")
		opts = append(opts, grpc.WithInsecure())
	}

	conn, err := grpc.Dial(cfg.thingsURL, opts...)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to connect to things service: %s", err))
		os.Exit(1)
	}
	return conn
}

func connectToRedis(redisURL, redisPass, redisDB string, logger logger.Logger) *redis.Client {
	db, err := strconv.Atoi(redisDB)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to connect to redis: %s", err))
		os.Exit(1)
	}

	return redis.NewClient(&redis.Options{
		Addr:     redisURL,
		Password: redisPass,
		DB:       db,
	})
}

func proxyMQTT(cfg config, logger logger.Logger, evt *mqtt.Event, errs chan error) {
	address := fmt.Sprintf("%s:%s", cfg.mqttHost, cfg.mqttPort)
	target := fmt.Sprintf("%s:%s", cfg.mqttTargetHost, cfg.mqttTargetPort)
	mp := mp.New(address, target, evt, logger)

	errs <- mp.Proxy()
}
func proxyWS(cfg config, logger logger.Logger, evt *mqtt.Event, errs chan error) {
	target := fmt.Sprintf("%s:%s", cfg.httpTargetHost, cfg.httpTargetPort)
	wp := ws.New(target, cfg.httpTargetPath, cfg.httpScheme, evt, logger)
	http.Handle("/mqtt", wp.Handler())

	p := fmt.Sprintf(":%s", cfg.httpPort)
	errs <- http.ListenAndServe(p, nil)
}
