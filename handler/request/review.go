package request

type Review struct {
	GoogleId  string   `json:"google_id" form:"google_id" query:"google_id"`
	Valuation float64 `json:"valuation" form:"valuation" query:"valuation"`
	Message   string  `json:"message" form:"message" query:"message"`
}
