package local

import (
	"time"
	"database/sql"
)

type Toilet struct {
	ID          int64
	name        string
	googleId    string
	lat         float64
	lng         float64
	geolocation string
	imagePath   string
	description string
	valuation   float64
	updatedAt   time.Time
}

func (toilet *Toilet) Insert(db sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	f := func(tx *sql.Tx) error {
		query := "INSERT INTO toilets(`name`,`google_id`,`lat`,`lng`,`geolocation`,`image_path`,`description`,`valuation`) VALUES (?,?,?,?,?,?,?,?)"
		_, err := tx.Exec(query, toilet.name, toilet.googleId, toilet.lat, toilet.lng, toilet.geolocation, toilet.imagePath, toilet.description, toilet.valuation)
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
