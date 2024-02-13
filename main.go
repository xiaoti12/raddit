package main

import (
	"errors"
	"fmt"
	"net/http"
	"raddit/config"
	"raddit/dao/mysql"
	"raddit/dao/redisdb"
	"raddit/logger"
	"raddit/pkg/snowflake"
	"raddit/routes"
)

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
	logger.Init(config.Conf.LogConfig)
	// init snowflake user id generator
	err = snowflake.Init(config.Conf.StartTime, config.Conf.MachineID)
	if err != nil {
		panic(err)
	}
	// 4. setup router engine
	r := routes.SetRouteEngine()
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
	quit := make(chan bool)
	<-quit
}
