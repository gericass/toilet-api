package response

import "time"

type Review struct {
	ToiletId  int64     `json:"toilet_id" form:"toilet_id" query:"toilet_id"`
	User      *User     `json:"user" form:"user" query:"user"`
	Valuation float64   `json:"valuation" form:"valuation" query:"valuation"`
	Message   string    `json:"message" form:"message" query:"message"`
	CreatedAt time.Time `json:"created_at" form:"created_at" query:"created_at"`
}
