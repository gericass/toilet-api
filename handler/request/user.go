package request

type User struct {
	GoogleId string `json:"google_id" form:"google_id" query:"google_id"`
	Name     string `json:"name" form:"name" query:"name"`
	Icon     string `json:"icon" form:"icon" query:"icon"`
}