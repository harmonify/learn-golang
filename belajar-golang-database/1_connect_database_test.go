package belajar_golang_database

import (
	"fmt"
	"testing"
)

func TestConnectDatabase(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	fmt.Println("Success connect to database")
}
