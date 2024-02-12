package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"raddit/config"
)

var db *sqlx.DB

func Init(cfg *config.MySQLConfig) error {
	//dsn := "root:root@tcp(127.0.0.1:3306)/raddit?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
	)
	var err error
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Println("connect to mysql database failed:", err)
		return err
	}
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	return nil
}
