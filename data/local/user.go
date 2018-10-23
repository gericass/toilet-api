package local

import (
	"time"
	"database/sql"
)

type User struct {
	ID        int64
	Name      string
	UID       string
	IconPath  string
	CreatedAt time.Time
}

func (user *User) Insert(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	f := func(tx *sql.Tx) error {
		query := "INSERT INTO users(`name`,`uid`,`icon_path`) VALUES (?,?,?)"
		_, err := tx.Exec(query, user.Name, user.UID, user.IconPath)
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

func (user *User) FindUserById(db *sql.DB) error {
	err := db.QueryRow("SELECT users(`id`, `name`, `uid`, `icon_path`, `created_at`) FROM users WHERE `id` = ?", user.ID).Scan(&user.ID, &user.Name, &user.UID, &user.IconPath, &user.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (user *User) FindUserByUID(db *sql.DB) error {
	err := db.QueryRow("SELECT users(`id`, `name`, `uid`, `icon_path`, `created_at`) FROM users WHERE `uid` = ?", user.UID).Scan(&user.ID, &user.Name, &user.UID, &user.IconPath, &user.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (user *User) Exists(db *sql.DB) (bool, error) {
	var count int
	err := db.QueryRow("SELECT count(*) FROM users WHERE `uid` = ?", user.UID).Scan(&count)
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

func (user *User) GetUserId(db *sql.DB) (int64, error) {
	var id int64
	err := db.QueryRow("SELECT `id` FROM users WHERE `uid` = ?", user.UID).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
