package response

type Toilets struct {
	Toilets []*Toilet `json:"toilets" form:"toilets" query:"toilets"`
	Status  string    `json:"status" form:"status" query:"status"`
}
