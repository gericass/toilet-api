package response

import "time"

type Toilet struct {
	Name        string `json:"name" form:"name" query:"name"`
	GoogleId    string `json:"google_id" form:"google_id" query:"google_id"`
	Geolocation string `json:"geolocation" form:"geolocation" query:"geolocation"`
	Image       string `json:"image" form:"image" query:"image"`
	Description string `json:"description" form:"description" query:"description"`
	Valuation   float64 `json:"valuation" form:"valuation" query:"valuation"`
	UpdatedAt   time.Time `json:"updated_at" form:"updated_at" query:"updated_at"`
}
