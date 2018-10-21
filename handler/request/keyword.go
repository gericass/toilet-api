package request

type Keyword struct {
	Keyword string `json:"keyword" form:"keyword" query:"keyword"`
}