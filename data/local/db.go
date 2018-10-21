package local

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func txHandler(tx *sql.Tx, f func(tx *sql.Tx) error) error {
	var err error
	err = f(tx)
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	return err
}

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:mysql@tcp(127.0.0.1:13306)/toilet")
	if err != nil {
		return nil, err
	}
	return db, nil
}
