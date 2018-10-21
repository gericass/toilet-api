package local

import (
	"time"
	"database/sql"
)

type User struct {
	ID        int64
	name      string
	googleId  string
	iconPath  string
	createdAt time.Time
}

func (user *User) Insert(db sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	f := func(tx *sql.Tx) error {
		query := "INSERT INTO users(`name`,`google_id`,`icon_path`) VALUES (?,?,?)"
		_, err := tx.Exec(query, user.name, user.googleId)
		if err != nil {
			return err
		}
		return nil
	}
	err = txHandler(tx, f)
	if err != nil {
		return err
	}
	return nil
}

func (user *User) Exists(db sql.DB) (bool, error) {
	var count int
	err := db.QueryRow("SELECT count(*) FROM users WHERE `google_id` = ?", user.googleId).Scan(&count)
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}
