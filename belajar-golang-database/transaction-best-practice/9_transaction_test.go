package transaction_best_practice

import (
	belajar_golang_database "belajar-golang-database"
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestTransactionBestPractice(t *testing.T) {
	defer func() {
		message := recover()
		if message != nil {
			fmt.Println("recovered panic message:", message)
		}
	}()

	Insert10Comments()
}

func Insert10Comments() {
	db := belajar_golang_database.GetConnection()
	defer db.Close()

	// This allows the transaction to run up to max 30 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	WithTransaction(ctx, db, func(tx Transaction) error {
		for i := 0; i < 10; i++ {
			email := "eko" + strconv.Itoa(i) + "@gmail.com"
			comment := "Komentar ke " + strconv.Itoa(i)

			// this could be improved by using prepared statement
			result, err := tx.Exec("INSERT INTO comments(email, comment) VALUES(?, ?)", email, comment)
			if err != nil {
				panic(err)
			}

			id, err := result.LastInsertId()
			if err != nil {
				panic(err)
			}

			if rand.Float32() >= 0.90 { // throw random error with 10% chance
				panic("example random panic that cause rollback")
			}

			fmt.Println("Comment Id", id)
		}

		return nil
	})
}
