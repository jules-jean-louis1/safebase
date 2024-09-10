package controllers

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

type DBParams struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
	DBType   string
}

func ConnectionTester(params *DBParams) (bool, error) {
	var dsn string

	if params != nil {
		switch params.DBType {
		case "postgres":
			dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
				params.Host, params.Port, params.Username, params.Password, params.DBName, params.SSLMode)
		case "mysql":
			dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
				params.Username, params.Password, params.Host, params.Port, params.DBName)
		default:
			return false, fmt.Errorf("unsupported database type: %s", params.DBType)
		}
	} else {
		return false, fmt.Errorf("either connStr or params must be provided")
	}

	db, err := sql.Open(params.DBType, dsn)
	if err != nil {
		return false, err
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		return false, err
	}

	return true, nil
}
