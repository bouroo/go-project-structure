package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bouroo/go-project-structure/api/auth"
	userRepository "github.com/bouroo/go-project-structure/api/user/repository"
	"github.com/bouroo/go-project-structure/datasources"
	"github.com/bouroo/go-project-structure/infrastructure"
	"github.com/bouroo/go-project-structure/infrastructure/config"
	"github.com/bouroo/go-project-structure/pkg/helper"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

var (
	RUN_ENV     string
	CONFIG_PATH string
)

func main() {
	var err error

	if RUN_ENV = os.Getenv("RUN_ENV"); len(RUN_ENV) == 0 {
		RUN_ENV = "local"
	}

	if CONFIG_PATH = os.Getenv("CONFIG_PATH"); len(CONFIG_PATH) == 0 {
		CONFIG_PATH = "./configs"
	}

	slog.Info("starting", "RUN_ENV", RUN_ENV)

	appConfig := config.NewAppConfig(CONFIG_PATH)
	if err = appConfig.LoadConfig(RUN_ENV); err != nil {
		log.Panic(err)
	} else {
		if err = appConfig.WatchConfig(); err != nil {
			log.Panic(err)
		}
	}

	if err = appConfig.CheckConfigs(config.KEYS_USED); err != nil {
		log.Panic(err)
	}
	datasources.AppConfig = appConfig.GetViper()

	// is debug mode with APP_DEBUG = true
	var handlerOptions slog.HandlerOptions
	logWriter := helper.LogWriter{
		ConsolePrint: true,
	}
	if appConfig.GetViper().GetBool("app.debug") {
		handlerOptions = slog.HandlerOptions{
			AddSource: true,
			Level:     slog.LevelDebug,
		}
	}

	// set logger target
	logger := slog.New(
		slog.NewTextHandler(
			logWriter,
			&handlerOptions,
		),
	)
	slog.SetDefault(logger)

	datasources.UserGRPCConn, err = infrastructure.NewGRPCConnectionPool(&infrastructure.GPRCClientConnConfig{
		TargetAddr:   fmt.Sprintf("%s:%d", datasources.AppConfig.GetString("service.user.grpc.host"), datasources.AppConfig.GetInt("service.user.grpc.port")),
		TLSConfig:    nil,
		MaxIdleConns: 100,
		DialTimeout:  10 * time.Second,
	})
	if err != nil {
		log.Panic(err)
	}

	datasources.DBConn, err = infrastructure.NewPostgresConn(infrastructure.PostgresOptions{
		Host:     appConfig.GetViper().GetString("db.postgres.host"),
		Port:     appConfig.GetViper().GetInt("db.postgres.port"),
		User:     appConfig.GetViper().GetString("db.postgres.user"),
		Password: appConfig.GetViper().GetString("db.postgres.password"),
		DBname:   appConfig.GetViper().GetString("db.postgres.database"),
		Debug:    appConfig.GetViper().GetBool("app.debug"),
	})
	if err != nil {
		log.Panic(err)
	}

	// Setup
	e := echo.New()
	e.Logger.SetLevel(log.INFO)
	e.JSONSerializer = &helper.CustomJSONSerializer{}

	// App middleware
	e.Use(echoMiddleware.Recover())
	if appConfig.GetViper().GetBool("app.debug") {
		e.Logger.SetLevel(log.DEBUG)
		e.Use(echoMiddleware.Logger())
	}
	e.Use(echoMiddleware.Secure())
	e.Use(echoMiddleware.RequestIDWithConfig(echoMiddleware.RequestIDConfig{
		Generator:        helper.CustomRequestIDGenerator,
		RequestIDHandler: helper.CustomRequestIDHandler,
	}))

	// Custom app handler
	e.HTTPErrorHandler = helper.CustomHTTPErrorHandler
	e.Validator = &helper.CustomValidator{Validator: validator.New()}

	if appConfig.GetViper().GetBool("app.automigrate") {
		err = userRepository.MigrateTable()
		if err != nil {
			log.Fatal(err)
		}
	}

	auth.RegisterRoute(e)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	// Start server
	listenFmt := `%s:%d`
	go func() {
		if err := e.Start(fmt.Sprintf(listenFmt, appConfig.GetViper().GetString("app.listen"), appConfig.GetViper().GetInt("app.port.http"))); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
	if datasources.DBConn != nil {
		if db, err := datasources.DBConn.DB(); err == nil {
			db.Close()
		}
	}
	if datasources.RedisConn != nil && datasources.RedisConn.Client != nil {
		datasources.RedisConn.Client.Close()
	}
	if datasources.UserGRPCConn != nil {
		datasources.UserGRPCConn.CloseAll()
	}
}
