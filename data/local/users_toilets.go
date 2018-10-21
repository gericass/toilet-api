package local

import (
	"time"
	"database/sql"
)

type UsersToilets struct {
	ID        int64
	UserId    int64
	ToiletId  int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (ut *UsersToilets) Insert(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	f := func(tx *sql.Tx) error {
		query := "INSERT INTO users_toilets(`user_id`,`toilet_id`) VALUES (?,?)"
		_, err := tx.Exec(query, ut.UserId, ut.ToiletId)
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

func (ut *UsersToilets) FindToiletsByUserId(db *sql.DB) ([]*UsersToilets, error) {
	rows, err := db.Query("SELECT `user_id`,`toilet_id`, `created_at`, `updated_at` FROM users_toilets WHERE `user_id` = ?", ut.UserId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var usersToilets []*UsersToilets

	for rows.Next() {
		t := &UsersToilets{}
		err := rows.Scan(&t.UserId, &t.ToiletId, &t.CreatedAt, &t.UpdatedAt)
		if err != nil {
			return nil, err
		}
		usersToilets = append(usersToilets, t)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return usersToilets, nil
}

func (ut *UsersToilets) RemoveUsersToilets(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	f := func(tx *sql.Tx) error {
		query := "DELETE FROM users_toilets WHERE `user_id` = ? AND `toilet_id` = ?"
		_, err := tx.Exec(query, ut.UserId, ut.ToiletId)
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
