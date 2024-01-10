package connectiondb

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

func NewConnectionDB(cfg Config) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// Set maximum number of open connections
	db.SetMaxOpenConns(10)

	// Set maximum number of idle connections
	db.SetMaxIdleConns(5)

	// Set maximum lifetime of a connection
	db.SetConnMaxLifetime(time.Hour)

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
