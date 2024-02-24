package main

import (
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"raddit/config"
	"raddit/dao/mysql"
	"raddit/dao/redisdb"
	"raddit/logger"
	"raddit/pkg/snowflake"
	"raddit/routes"
	"syscall"
	"time"
)

// @title Raddit
// @version 1.0
// @description This is a simple forum service.
// @host 127.0.0.1:8898
// @BasePath /api
func main() {
	var err error
	// 1. init config
	err = config.Init()
	if err != nil {
		panic(err)
	}
	// 2. init database
	err = mysql.Init(config.Conf.MySQLConfig)
	if err != nil {
		panic(err)
	}

	err = redisdb.Init(config.Conf.RedisConfig)
	if err != nil {
		panic(err)
	}
	// 3. init logger
	logger.Init(config.Conf.LogConfig, config.Conf.Mode)
	// init snowflake user id generator
	err = snowflake.Init(config.Conf.StartTime, config.Conf.MachineID)
	if err != nil {
		panic(err)
	}
	// 4. setup router engine
	r := routes.SetRouteEngine(config.Conf.Mode)
	// 5. start service
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Conf.Port),
		Handler: r,
	}
	go func() {
		err := server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}()
	// 6. graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = server.Shutdown(ctx)
	if err != nil {
		zap.L().Fatal("server shutdown error", zap.Error(err))
	}
	fmt.Println("server exit")
}
