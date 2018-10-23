package response

type User struct {
	Name     string `json:"name" form:"name" query:"name"`
	UID      string `json:"uid" form:"uid" query:"uid"`
	IconPath string `json:"icon" form:"icon" query:"icon"`
}
