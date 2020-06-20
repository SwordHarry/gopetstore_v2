package domain

type Category struct {
	CategoryId  string `db:"catid"`
	Name        string `db:"name"`
	Description string `db:"descn"`
}
