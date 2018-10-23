package response

type Reviews struct {
	Reviews []*Review `json:"reviews" form:"reviews" query:"reviews"`
	Status  string    `json:"status" form:"status" query:"status"`
}
