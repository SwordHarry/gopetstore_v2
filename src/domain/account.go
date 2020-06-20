package domain

import "encoding/gob"

// 登录用户
// 由 signon, profile, account 三张表组成
type Account struct {
	// account
	UserName  string `db:"userid"`
	Email     string `db:"email"`
	FirstName string `db:"firstname"`
	LastName  string `db:"lastname"`
	Status    string `db:"status"`
	Address1  string `db:"addr1"`
	Address2  string `db:"addr2"`
	City      string `db:"city"`
	State     string `db:"state"`
	Zip       string `db:"zip"`
	Country   string `db:"country"`
	Phone     string `db:"phone"`
	// signon
	Password string `db:"password"`
	// profile
	FavouriteCategoryId string `db:"favcategory"`
	LanguagePreference  string `db:"langpref"`
	ListOption          bool   `db:"mylistopt"`
	BannerOption        bool   `db:"banneropt"`
	// banner
	BannerName string `db:"bannername"`
}

func init() {
	gob.Register(&Account{})
}
