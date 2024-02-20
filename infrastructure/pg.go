package infrastructure

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresOptions struct {
	Host         string
	Port         int
	User         string
	Password     string
	DBname       string
	SSLMode      string
	Timezone     string
	MaxIdle      int
	MaxOpen      int
	ConnLifetime time.Duration
	Debug        bool
}

func (opt *PostgresOptions) ApplyDefault() *PostgresOptions {
	if len(opt.Host) == 0 {
		opt.Host = "localhost"
	}
	if opt.Port == 0 {
		opt.Port = 5432
	}
	if len(opt.User) == 0 {
		opt.User = "postgres"
	}
	if len(opt.Password) == 0 {
		opt.Password = "postgres"
	}
	if len(opt.DBname) == 0 {
		opt.DBname = "postgres"
	}
	if len(opt.SSLMode) == 0 {
		opt.SSLMode = "disable"
	}
	if len(opt.Timezone) == 0 {
		opt.Timezone = "Asia/Bangkok"
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

func NewPostgresConn(opts PostgresOptions) (db *gorm.DB, err error) {

	opts.ApplyDefault()

	dsn := "host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s"
	dsn = fmt.Sprintf(dsn, opts.Host, opts.User, opts.Password, opts.DBname, opts.Port, opts.SSLMode, opts.Timezone)

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

	if opts.Debug {
		db = db.Debug()
	}

	return
}
