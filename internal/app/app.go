package app

import (
	"fmt"
	"goevents/config"
	"goevents/internal/interface/delivery/http"
	"goevents/internal/interface/delivery/http/v1/controllers"
	"goevents/internal/wire"
	"goevents/pkg/httpserver"
	"goevents/pkg/logger"
	"goevents/pkg/mysql"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(cfg *config.Config) {
	log := logger.New(cfg.Log.Level)

	// Initialize MYSQL
	mysqlUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.MYSQL.Username, cfg.MYSQL.Password, cfg.MYSQL.Host, cfg.MYSQL.Port, cfg.MYSQL.Database)
	mysql, err := mysql.New(mysqlUrl, mysql.SetMaxIdleConns(cfg.MYSQL.PoolMax), mysql.SetMaxOpenConns(cfg.MYSQL.PoolMax), mysql.SetConnMaxLifetime(time.Duration(cfg.MYSQL.PoolMax)*time.Second))
	if err != nil {
		panic(err)
	}
	defer mysql.Close()

	httpserver := httpserver.New(httpserver.Port(cfg.HTTP.Port))
	httpserver.Start()

	dbPingFunc := func() error {
		return mysql.DB.Ping()
	}

	healthCheck := controllers.NewHealthController(dbPingFunc)

	http.NewRouter(httpserver.Router, http.RouterOption{
		EventController:  wire.InitializeEventController(mysql.DB),
		HealthController: healthCheck,
	})

	// Wait for interrupt signal to gracefully shutdown the server with
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-quit:
		log.Info("internal/app - Run - signal: %s", s.String())
	case err := <-httpserver.Notify():
		log.Error(fmt.Errorf("internal/app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpserver.Shutdown()
	if err != nil {
		log.Error(fmt.Errorf("internal/app - Run - httpServer.Shutdown: %w", err))
	}
}
