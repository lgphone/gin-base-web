package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"
	"younghe/config"
	"younghe/lib/loggerx"
	"younghe/lib/migrate"
	"younghe/lib/mysqlx"
	"younghe/lib/redisx"
	"younghe/router"
)

var confPath string
var mode string

func setup() {
	config.Setup(confPath)
	loggerx.Setup()
	mysqlx.Setup()
	redisx.Setup()
}

func main() {
	// cmd args
	flag.StringVar(&confPath, "c", "./config.yaml", "配置文件路径.")
	flag.StringVar(&mode, "mode", "server", "运行模式.")
	flag.Parse()

	// setup
	setup()

	// migrate
	if mode == "migrate" {
		loggerx.Logger.Info("migrate tables...")
		migrate.Migrate()
		return
	}

	// init router
	engine := router.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", config.Config.Server.Port),
		Handler:        engine,
		MaxHeaderBytes: 1 << 20,
	}

	// start http server
	go func() {
		loggerx.Logger.Info("start http server on:", config.Config.Server.Port)
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			os.Exit(1)
		}
	}()

	// wait for signal to quit
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	// close all db
	loggerx.Logger.Info("close db ...")
	_ = mysqlx.DB.Close()
	_ = redisx.Redis.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// shutdown http server
	loggerx.Logger.Info("shutdown http server ...")
	if err := s.Shutdown(ctx); err != nil {
		os.Exit(1)
	}
}
