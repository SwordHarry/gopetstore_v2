package service

import (
	"gopetstore_v2/src/domain"
	"gopetstore_v2/src/persistence"
)

// get account by user name
func GetAccountByUserName(userName string) (*domain.Account, error) {
	return persistence.GetAccountByUserName(userName)
}

// get account by user name and password
func GetAccountByUserNameAndPassword(userName string, password string) (*domain.Account, error) {
	return persistence.GetAccountByUserNameAndPassword(userName, password)
}

// insert account
func InsertAccount(a *domain.Account) error {
	return persistence.InsertAccount(a)
}

// update account
func UpdateAccount(a *domain.Account) error {
	return persistence.UpdateAccount(a)
}
