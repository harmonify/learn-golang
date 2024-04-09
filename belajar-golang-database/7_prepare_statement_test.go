package belajar_golang_database

import (
	"context"
	"fmt"
	"testing"
)

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	stmt, err := db.PrepareContext(ctx, "INSERT INTO comments(email, comment) VALUES (?, ?);")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	for i := 1; i <= 10; i++ {
		email := fmt.Sprintf("wendy+%d@gmail.com", i)
		comment := fmt.Sprintf("Comment %d", i)

		result, err := stmt.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}

		lastInsertId, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Comment id", lastInsertId)
	}
}
