package database

import (
	"fmt"
	"go-gin/config"
	"go-gin/core/log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

type Database struct {}

var (
	conn *xorm.Engine
	once sync.Once
	err  error
)

func Conn() *xorm.Engine {
	once.Do(func() {
		new()
	})
	return conn
}

func new() {

	switch config.Database["drive"].(string) {
	case "mysql":
		conn, err = mysqlConnection()
	case "postgres":
		conn, err = postgresConnection()
	default:
		conn, err = mysqlConnection()
	}

	if err != nil {
		log.Error(err)
		return
	}
	
	if maxIdleConns := config.Database["maxIdleConns"]; maxIdleConns != nil {
		conn.SetMaxIdleConns(maxIdleConns.(int))
	}

	if maxOpenConns := config.Database["maxOpenConns"]; maxOpenConns != nil {
		conn.SetMaxOpenConns(maxOpenConns.(int))
	}

	conn.SetMapper(names.SameMapper{})

	return
}

// mysql connects to the database
// mysql 连接数据库
func mysqlConnection() (*xorm.Engine, error) {

	host := config.Database["host"]
	port := config.Database["port"]
	dbname := config.Database["dbname"]
	username := config.Database["username"]
	password := config.Database["password"]

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		username, password, host, port, dbname,
	)

	return xorm.NewEngine("mysql", dsn)
}

// postgres connects to the database
// postgres 连接数据库
func postgresConnection() (*xorm.Engine, error) {

	host := config.Database["host"]
	port := config.Database["port"]
	dbname := config.Database["dbname"]
	username := config.Database["username"]
	password := config.Database["password"]

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		host, username, password, dbname, port,
	)

	return xorm.NewEngine("postgres", dsn)
}
