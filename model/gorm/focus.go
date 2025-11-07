package gorm

type Focus struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Pic      string `json:"pic"`
	Link     string `json:"link"`
	Position int    `json:"position"`
}

func (Focus) TableName() string {
	return "focus"
}
