package mysql

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	Host     string `yml:"host"`
	Port     int    `yaml:"port"`
	Username string `yml:"username"`
	Password string `yml:"password"`
	DBName   string `yml:"dbname"`
}

type MySQLDB struct {
	config Config
	db     *sql.DB
}

func (m *MySQLDB) Conn() *sql.DB {
	return m.db
}

func New(cfg Config) *MySQLDB {

	uri := fmt.Sprintf("%s:%s@(%s:%d)/%s?parseTime=true", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)

	db, err := sql.Open("mysql", uri)
	if err != nil {
		panic(err)
	}

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return &MySQLDB{config: cfg, db: db}
}
