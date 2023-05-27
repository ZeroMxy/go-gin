package database

import (
	"go-gin/config"
	"go-gin/core/log"
	"strconv"
	"sync"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	// _ "github.com/mattn/go-oci8"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

type Database struct {}

var (
	conn *xorm.Engine
	once sync.Once
	err  error
)

func Conn () *xorm.Engine {
	
	once.Do(func() {
		new()
	})
	return conn
}

func new () {

	switch config.Database["drive"] {
	case "mysql":
		conn, err = mysqlConn()
	case "postgres":
		conn, err = postgresConn()
	case "mssql":
		conn, err = mssqlConn()
	// case "oracle":
	// 	conn, err = oracleConn()
	default:
		conn, err = mysqlConn()
	}

	if err != nil {
		log.Error(err)
		return
	}

	maxIdleConns, err := strconv.Atoi(config.Database["maxIdleConns"])
	if err != nil {
		conn.SetMaxIdleConns(maxIdleConns)
	}

	maxOpenConns, err := strconv.Atoi(config.Database["maxOpenConns"])
	if err != nil {
		conn.SetMaxOpenConns(maxOpenConns)
	}

	connMaxLifetime, err := time.ParseDuration(config.Database["connMaxLifetime"])
	if err != nil {
		conn.SetConnMaxLifetime(connMaxLifetime)
	}

	conn.SetMapper(names.SameMapper{})

	return
}

// mysql connects to the database
// mysql 连接数据库
func mysqlConn () (*xorm.Engine, error) {

	host := config.Database["host"]
	port := config.Database["port"]
	dbname := config.Database["dbname"]
	username := config.Database["username"]
	password := config.Database["password"]

	// dsn := "username:password@tcp(host:port)/dbname?charset=utf8&parseTime=True&loc=Local"
	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8&parseTime=True&loc=Local"

	return xorm.NewEngine("mysql", dsn)
}

// postgres connects to the database
// postgres 连接数据库
func postgresConn () (*xorm.Engine, error) {

	host := config.Database["host"]
	port := config.Database["port"]
	dbname := config.Database["dbname"]
	username := config.Database["username"]
	password := config.Database["password"]

	// dsn := "host=host user=username password=password dbname=dbname port=port sslmode=disable"
	dsn := "host=" + host + " user=" + username + " password=" + password + " dbname=" + dbname + " port=" + port + "sslmode=disable"

	return xorm.NewEngine("postgres", dsn)
}

// mssql connects to the database
// mssql 连接数据库
func mssqlConn () (*xorm.Engine, error) {

	host := config.Database["host"]
	port := config.Database["port"]
	dbname := config.Database["dbname"]
	username := config.Database["username"]
	password := config.Database["password"]

	// dsn := "server=host;port=port;database=dbname;user id=username;password=password"
	dsn := "server=" + host + ";port=" + port + ";database=" + dbname+ ";user id=" + username + ";password=" + password

	return xorm.NewEngine("mssql", dsn)
}

// oracle connects to the database
// oracle 连接数据库
func oracleConn () (*xorm.Engine, error) {

	host := config.Database["host"]
	port := config.Database["port"]
	dbname := config.Database["dbname"]
	username := config.Database["username"]
	password := config.Database["password"]

	// dsn := "username/password@host:port/dbname"
	dsn := username + "/" + password + "@" + host + ":" + port + "/" + dbname

	return xorm.NewEngine("oci8", dsn)
}
