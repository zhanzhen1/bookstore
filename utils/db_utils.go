package utils

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var (
	DB *sql.DB
)

func InitDB() error {
	dsn := os.Getenv("MYSQL_DSN")
	if dsn == "" {
		return errors.New("env MYSQL_DSN must set")
	}
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("open database: %w", err)
	}
	if err := DB.Ping(); err != nil {
		return fmt.Errorf("ping database: %w", err)
	}
	return nil
}
