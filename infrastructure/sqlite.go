package infrastructure

import (
	"fmt"
	"net/url"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SQLiteOptions struct {
	Path         string
	Params       url.Values
	MaxIdle      int
	MaxOpen      int
	ConnLifetime time.Duration
}

func (opt *SQLiteOptions) ApplyDefault() *SQLiteOptions {
	if len(opt.Path) == 0 {
		opt.Path = "file::memory:"
	}
	if len(opt.Params) == 0 {
		opt.Params = url.Values{
			"cache": {"shared"},
		}
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

func NewSQLiteConn(opts SQLiteOptions) (db *gorm.DB, err error) {

	opts.ApplyDefault()

	dsn := "%s%s"
	dsn = fmt.Sprintf(dsn, opts.Path, opts.Params.Encode())

	db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
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
