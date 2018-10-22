package request

type User struct {
	UID  string `json:"uid" form:"uid" query:"uid"`
	Name string `json:"name" form:"name" query:"name"`
	Icon string `json:"icon" form:"icon" query:"icon"`
}
