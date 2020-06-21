package domain

type Category struct {
	CategoryId  string `db:"catid"`
	Name        string `db:"catname"`
	Description string `db:"descn"`
}
