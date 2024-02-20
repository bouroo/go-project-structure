package infrastructure

import (
	"fmt"
	"net/url"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type MySQLOptions struct {
	Host         string
	Port         int
	User         string
	Password     string
	DBname       string
	ParseTime    bool
	Charset      string
	Locale       string
	MaxIdle      int
	MaxOpen      int
	ConnLifetime time.Duration
}

func (opt *MySQLOptions) ApplyDefault() *MySQLOptions {
	if len(opt.Host) == 0 {
		opt.Host = "localhost"
	}
	if opt.Port == 0 {
		opt.Port = 5432
	}
	if len(opt.User) == 0 {
		opt.User = "mariadb"
	}
	if len(opt.Password) == 0 {
		opt.Password = "mariadb"
	}
	if len(opt.DBname) == 0 {
		opt.DBname = "mariadb"
	}
	if len(opt.Charset) == 0 {
		opt.Charset = "utf8mb4,utf8"
	}
	if len(opt.Locale) == 0 {
		opt.Locale = "Asia/Bangkok"
	}
	if opt.MaxIdle == 0 {
		opt.MaxIdle = 10
	}
	if opt.MaxOpen == 0 {
		opt.MaxOpen = 100
	}
	if opt.ConnLifetime == 0 {
		opt.ConnLifetime = time.Hour
	}
	return opt
}

func NewMySQLConn(opts MySQLOptions) (db *gorm.DB, err error) {

	opts.ApplyDefault()

	opts.Locale = url.QueryEscape(opts.Locale)

	dsn := "%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=%s"
	dsn = fmt.Sprintf(dsn, opts.User, opts.Password, opts.Host, opts.Port, opts.DBname, opts.Charset, opts.ParseTime, opts.Locale)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}

	sqlDB, err := db.DB()
	if err != nil {
		return
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(opts.MaxIdle)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(opts.MaxOpen)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(opts.ConnLifetime)
	return
}
