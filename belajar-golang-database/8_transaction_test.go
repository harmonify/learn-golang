package belajar_golang_database

import (
	"context"
	"fmt"
	"strconv"
	"testing"
)

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	// https://stephenn.com/2023/08/mastering-database-transactions-in-go-strategies-and-best-practices/
	// Always do a rollback at the end of the function.
	// If committed this is a noop.
	// If there was an error this closes the transaction.
	defer func() {
		err := tx.Rollback()
		if err == nil {
			fmt.Println("Rolled back")
		}
	}()

	script := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	// do transaction
	for i := 0; i < 10; i++ {
		email := "eko" + strconv.Itoa(i) + "@gmail.com"
		comment := "Komentar ke " + strconv.Itoa(i)

		result, err := tx.ExecContext(ctx, script, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Comment Id ", id)
	}

	// panic("example panic to cause rollback")

	err = tx.Commit()
	if err != nil {
		panic(err)
	}
	fmt.Println("Commited")
}
