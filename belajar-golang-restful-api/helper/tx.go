package helper

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		PanicIfError(tx.Rollback())
		panic(err)
	} else {
		PanicIfError(tx.Commit())
	}
}
