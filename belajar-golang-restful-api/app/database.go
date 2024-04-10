package app

import (
	"belajar-golang-restful-api/helper"
	"database/sql"
	"fmt"
	"time"
)

func ConstructDataSourceName(
	host string,
	port int,
	database string,
	username string,
	password string,
) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", username, password, host, port, database)
}

func NewDB() *sql.DB {
	dataSourceName := ConstructDataSourceName("localhost", 3306, "belajar_golang_restful_api", "root", "")

	db, err := sql.Open("mysql", dataSourceName)
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
