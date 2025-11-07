package gorm

type Article struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Description string `json:"description"`
	Picture     string `json:"picture"`
}

func (Article) TableName() string {
	return "article"
}
