// https://gist.github.com/pseudomuto/0900a7a3605470760579752fcf0fc2b7
package transaction_best_practice

import (
	"context"
	"database/sql"
	"fmt"
)

// Transaction is an interface that models the standard transaction in
// `database/sql`.
//
// To ensure `TxFn` funcs cannot commit or rollback a transaction (which is
// handled by `WithTransaction`), those methods are not included here.
type Transaction interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	Stmt(stmt *sql.Stmt) *sql.Stmt
}

// A Txfn is a function that will be called with an initialized `Transaction` object
// that can be used for executing statements and queries against a database.
type TxFn func(Transaction) error

// WithTransaction creates a new transaction and handles rollback/commit based on the
// error object returned by the `TxFn`
func WithTransaction(ctx context.Context, db *sql.DB, fn TxFn) (err error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return
	}

	defer func() {
		if p := recover(); p != nil {
			// a panic occurred, rollback and repanic
			err2 := tx.Rollback()
			if err2 == nil {
				fmt.Println("Rolled back and repanic")
			}
			panic(p)
		} else if err != nil {
			// something went wrong, rollback
			err2 := tx.Rollback()
			if err2 == nil {
				fmt.Println("Rolled back")
			}
		} else {
			// all good, commit
			err = tx.Commit()
			if err == nil {
				fmt.Println("Commited")
			}
		}
	}()

	err = fn(tx)
	return err
}
