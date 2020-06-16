package domain

import "encoding/gob"

// 登录用户
// 由 signon, profile, account 三张表组成
type Account struct {
	UserName            string
	Password            string
	Email               string
	FirstName           string
	LastName            string
	Status              string
	Address1            string
	Address2            string
	City                string
	State               string
	Zip                 string
	Country             string
	Phone               string
	FavouriteCategoryId string
	LanguagePreference  string
	ListOption          bool
	BannerOption        bool
	BannerName          string
}

func init() {
	gob.Register(&Account{})
}
