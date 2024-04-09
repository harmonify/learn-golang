package belajar_golang_database

import (
	"context"
	"fmt"
	"testing"
)

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	email := "wendy@gmail.com"
	comment := "A comment"

	query := "INSERT INTO comments(email, comment) VALUES (?, ?);"
	rows, err := db.ExecContext(ctx, query, email, comment)
	if err != nil {
		panic(err)
	}

	lastInsertId, err := rows.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new comment with id", lastInsertId, "to database")
}
