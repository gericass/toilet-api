package local

import (
	"time"
	"database/sql"
)

type Toilet struct {
	ID          int64
	Name        string
	GoogleId    string
	Lat         float64
	Lng         float64
	Geolocation string
	ImagePath   string
	Description string
	Valuation   float64
	UpdatedAt   time.Time
}

func (toilet *Toilet) Insert(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	f := func(tx *sql.Tx) error {
		query := "INSERT INTO toilets(`name`,`google_id`,`lat`,`lng`,`geolocation`,`image_path`,`description`,`valuation`) VALUES (?,?,?,?,?,?,?,?)"
		_, err := tx.Exec(query, toilet.Name, toilet.GoogleId, toilet.Lat, toilet.Lng, toilet.Geolocation, toilet.ImagePath, toilet.Description, toilet.Valuation)
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

func (toilet *Toilet) Exists(db *sql.DB) (bool, error) {
	var count int
	err := db.QueryRow("SELECT count(*) FROM toilets WHERE `google_id` = ?", toilet.GoogleId).Scan(&count)
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

func (toilet *Toilet) FindToiletByGoogleId(db *sql.DB) error {
	err := db.QueryRow("SELECT `id`,`name`, `lat`, `lng`, `image_path`, `description`, `valuation`, `updated_at` FROM toilets WHERE `google_id` = ?", toilet.GoogleId).Scan(&toilet.ID, &toilet.Name, &toilet.Lat, &toilet.Lng, &toilet.ImagePath, &toilet.Description, &toilet.Valuation, &toilet.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (toilet *Toilet) FindToiletById(db *sql.DB) error {
	err := db.QueryRow("SELECT `name`, `google_id`,`lat`, `lng`, `image_path`, `description`, `valuation`, `updated_at` FROM toilets WHERE `id` = ?", toilet.ID).Scan(&toilet.Name, &toilet.GoogleId, &toilet.Lat, &toilet.Lng, &toilet.ImagePath, &toilet.Description, &toilet.Valuation, &toilet.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (toilet *Toilet) GetToiletId(db *sql.DB) (int64, error) {
	var id int64
	err := db.QueryRow("SELECT `id` FROM toilets WHERE `google_id` = ?", toilet.GoogleId).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (toilet *Toilet) UpdateValuation(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	f := func(tx *sql.Tx) error {
		query := "UPDATE toilets SET `valuation` = ? WHERE `toilet_id` = ?"
		_, err := tx.Exec(query, toilet.Valuation, toilet.ID)
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
