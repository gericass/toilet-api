package local

import (
	"time"
	"database/sql"
)

type Review struct {
	ID        int64
	ToiletId  int64
	UserId    int64
	Valuation float64
	Message   string
	CreatedAt time.Time
}

func (review *Review) Insert(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	f := func(tx *sql.Tx) error {
		query := "INSERT INTO reviews(`toilet_id`,`user_id`,`valuation`,`message`) VALUES (?,?,?,?)"
		_, err := tx.Exec(query, review.ToiletId, review.UserId, review.Valuation, review.Message)
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

func (review *Review) Exists(db *sql.DB) (bool, error) {
	var count int
	err := db.QueryRow("SELECT count(*) FROM reviews WHERE `toilet_id` = ? AND `user_id` = ?", review.ToiletId, review.UserId).Scan(&count)
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

func (review *Review) ExistsByToiletId(db *sql.DB) (bool, error) {
	var count int
	err := db.QueryRow("SELECT count(*) FROM reviews WHERE `toilet_id` = ?", review.ToiletId).Scan(&count)
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

func (review *Review) ExistsByUserId(db *sql.DB) (bool, error) {
	var count int
	err := db.QueryRow("SELECT count(*) FROM reviews WHERE `user_id` = ?", review.UserId).Scan(&count)
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

func (review *Review) FindReviewsByToiletId(db *sql.DB) ([]*Review, error) {
	rows, err := db.Query("SELECT `id`,`toilet_id`, `user_id`, `valuation`, `message`, `created_at` FROM reviews WHERE `toilet_id` = ? LIMIT 50", review.ToiletId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reviews []*Review

	for rows.Next() {
		r := &Review{}
		err := rows.Scan(&r.ID, &r.ToiletId, &r.UserId, &r.Valuation, &r.Message, &r.CreatedAt)
		if err != nil {
			return nil, err
		}
		reviews = append(reviews, r)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return reviews, nil
}

func (review *Review) FindReviewsByUserId(db *sql.DB) ([]*Review, error) {
	rows, err := db.Query("SELECT `id`,`toilet_id`, `user_id`, `valuation`, `message`, `created_at` FROM reviews WHERE `user_id` = ? LIMIT 50", review.UserId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reviews []*Review

	for rows.Next() {
		r := &Review{}
		err := rows.Scan(&r.ID, &r.ToiletId, &r.UserId, &r.Valuation, &r.Message, &r.CreatedAt)
		if err != nil {
			return nil, err
		}
		reviews = append(reviews, r)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return reviews, nil
}

func (review *Review) DeleteReview(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	f := func(tx *sql.Tx) error {
		query := "DELETE FROM reviews WHERE `user_id` = ? AND `toilet_id` = ?"
		_, err := tx.Exec(query, review.UserId, review.ToiletId)
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
