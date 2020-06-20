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
	err := persistence.InsertAccount(a)
	if err != nil {
		return err
	}
	err = persistence.InsertProfile(a)
	if err != nil {
		return err
	}
	err = persistence.InsertSignOn(a)
	if err != nil {
		return err
	}
	return nil
}

// update account
func UpdateAccount(a *domain.Account) error {
	err := persistence.UpdateAccount(a)
	if err != nil {
		return err
	}
	err = persistence.UpdateProfile(a)
	if err != nil {
		return err
	}
	err = persistence.UpdateSignOn(a)
	if err != nil {
		return err
	}
	return nil
}
