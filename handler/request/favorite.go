package request

type Favorite struct {
	GoogleId string `json:"google_id" form:"google_id" query:"google_id"`
}
