package belajar_golang_database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
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

func GetConnection() *sql.DB {
	dataSourceName := ConstructDataSourceName("localhost", 3306, "belajar_golang_database", "root", "p@55w0rd")
	fmt.Println(dataSourceName)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(200)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
