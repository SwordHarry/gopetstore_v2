package domain

import "encoding/gob"

type Product struct {
	ProductId   string `db:"productid"`
	CategoryId  string `db:"category"`
	Name        string `db:"name"`
	Description string `db:"descn"`
}

func init() {
	gob.Register(&Product{})
}
