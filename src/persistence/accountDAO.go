package persistence

import (
	"errors"
	"gopetstore_v2/src/domain"
	"gopetstore_v2/src/util"
)

const getAccountByUsernameSQL = `SELECT SIGNON.USERNAME,ACCOUNT.EMAIL,ACCOUNT.FIRSTNAME,ACCOUNT.LASTNAME,ACCOUNT.STATUS,ACCOUNT.ADDR1 AS address1,
ACCOUNT.ADDR2 AS address2,ACCOUNT.CITY,ACCOUNT.STATE,ACCOUNT.ZIP,ACCOUNT.COUNTRY,ACCOUNT.PHONE,PROFILE.LANGPREF AS languagePreference,
PROFILE.FAVCATEGORY AS favouriteCategoryId,PROFILE.MYLISTOPT AS listOption,PROFILE.BANNEROPT AS bannerOption,BANNERDATA.BANNERNAME 
FROM ACCOUNT, PROFILE, SIGNON, BANNERDATA
WHERE ACCOUNT.USERID = ? AND SIGNON.USERNAME = ACCOUNT.USERID AND PROFILE.USERID = ACCOUNT.USERID AND PROFILE.FAVCATEGORY = BANNERDATA.FAVCATEGORY`

// get account by userName and password from signOn, account, bannerData
const getAccountByUsernameAndPasswordSQL = `SELECT SIGNON.USERNAME,ACCOUNT.EMAIL,ACCOUNT.FIRSTNAME,ACCOUNT.LASTNAME,
ACCOUNT.STATUS,ACCOUNT.ADDR1 AS address1,ACCOUNT.ADDR2 AS address2,ACCOUNT.CITY,ACCOUNT.STATE,ACCOUNT.ZIP,
ACCOUNT.COUNTRY,ACCOUNT.PHONE,PROFILE.LANGPREF AS languagePreference,PROFILE.FAVCATEGORY AS favouriteCategoryId,
PROFILE.MYLISTOPT AS listOption,PROFILE.BANNEROPT AS bannerOption,BANNERDATA.BANNERNAME FROM ACCOUNT, PROFILE, SIGNON, BANNERDATA 
WHERE ACCOUNT.USERID = ? AND SIGNON.PASSWORD = ? AND SIGNON.USERNAME = ACCOUNT.USERID AND 
PROFILE.USERID = ACCOUNT.USERID AND PROFILE.FAVCATEGORY = BANNERDATA.FAVCATEGORY`

// Insert
// insert account from account
const insertAccountSQL = `INSERT INTO ACCOUNT (EMAIL, FIRSTNAME, LASTNAME, STATUS, ADDR1, ADDR2, CITY, STATE, ZIP, COUNTRY, PHONE, USERID) 
VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

// insert profile from profile
const insertProfileSQL = `INSERT INTO PROFILE (LANGPREF, FAVCATEGORY, USERID, mylistopt, banneropt) VALUES (?, ?, ?, ?, ?)`

// insert username and password from signOn
const insertSigOnSQL = `INSERT INTO SIGNON (USERNAME,PASSWORD) VALUES (?, ?)`

// Update
// update account from account
const updateAccountSQL = `UPDATE ACCOUNT SET EMAIL = ?,FIRSTNAME = ?,LASTNAME = ?,STATUS = ?,ADDR1 = ?,ADDR2 = ?,
CITY = ?,STATE = ?,ZIP = ?,COUNTRY = ?,PHONE = ? WHERE USERID = ?`

// update profile from profile
const updateProfileSQL = `UPDATE PROFILE SET LANGPREF = ?, FAVCATEGORY = ?,mylistopt = ?,banneropt = ? WHERE USERID = ?`

// update password by userName from signOn
const updateSigOnSQL = `UPDATE SIGNON SET PASSWORD = ? WHERE USERNAME = ?`

// query
// get account by user name
func GetAccountByUserName(userName string) (*domain.Account, error) {
	d, err := util.GetConnection()
	defer func() {
		_ = d.Close()
	}()
	if err != nil {
		return nil, err
	}
	a := new(domain.Account)
	// 这里涉及 account, signon, profile, bannerdata 四个表的连接
	// 使用 sqlx 的 Get api， 减少代码量
	err = d.Get(a, getAccountByUsernameSQL, userName)
	if err != nil {
		return nil, err
	}
	return a, nil
}

// get account
func GetAccountByUserNameAndPassword(userName string, password string) (*domain.Account, error) {
	d, err := util.GetConnection()
	defer func() {
		_ = d.Close()
	}()
	if err != nil {
		return nil, err
	}
	a := new(domain.Account)
	// 这里涉及 account, signon, profile, bannerdata 四个表的连接
	err = d.Get(a, getAccountByUsernameAndPasswordSQL, userName, password)
	if err != nil {
		return nil, err
	}
	return a, nil
}

// insert
// 插入 account 数据表
func InsertAccount(a *domain.Account) error {
	d, err := util.GetConnection()
	defer func() {
		_ = d.Close()
	}()
	if err != nil {
		return err
	}
	r, err := d.Exec(insertAccountSQL, a.Email, a.FirstName, a.LastName, a.Status,
		a.Address1, a.Address2, a.City, a.State, a.Country, a.Phone, a.UserName)
	if err != nil {
		return err
	}
	row, err := r.RowsAffected()
	if err != nil {
		return err
	}
	if row == 0 {
		return errors.New("can not insert account by this user name:" + a.UserName)
	}
	return nil
}

// 插入 profile 数据表
func InsertProfile(a *domain.Account) error {
	d, err := util.GetConnection()
	defer func() {
		_ = d.Close()
	}()
	if err != nil {
		return err
	}
	r, err := d.Exec(insertProfileSQL, a.LanguagePreference, a.FavouriteCategoryId, a.UserName, a.ListOption, a.BannerOption)
	if err != nil {
		return err
	}
	row, err := r.RowsAffected()
	if err != nil {
		return err
	}
	if row == 0 {
		return errors.New("can not insert profile by this user name:" + a.UserName)
	}
	return nil
}

// 插入 signon 数据表
func InsertSignOn(a *domain.Account) error {
	d, err := util.GetConnection()
	defer func() {
		_ = d.Close()
	}()
	if err != nil {
		return err
	}
	r, err := d.Exec(insertSigOnSQL, a.UserName, a.Password)
	if err != nil {
		return err
	}
	row, err := r.RowsAffected()
	if err != nil {
		return err
	}
	if row == 0 {
		return errors.New("can not insert sign on by this user name:" + a.UserName)
	}
	return nil
}

// update
func UpdateAccount(a *domain.Account) error {
	d, err := util.GetConnection()
	if err != nil {
		return err
	}
	r, err := d.Exec(updateAccountSQL, a.Email, a.FirstName, a.LastName, a.Status, a.Address1, a.Address2,
		a.City, a.State, a.Zip, a.Country, a.Phone, a.UserName)
	if err != nil {
		return err
	}
	row, err := r.RowsAffected()
	if err != nil {
		return err
	}
	if row == 0 {
		return errors.New("can not update account by this user name:" + a.UserName)
	}
	return nil
}

func UpdateProfile(a *domain.Account) error {
	d, err := util.GetConnection()
	if err != nil {
		return err
	}
	r, err := d.Exec(updateProfileSQL, a.LanguagePreference, a.FavouriteCategoryId, a.ListOption, a.BannerOption, a.UserName)
	if err != nil {
		return err
	}
	row, err := r.RowsAffected()
	if err != nil {
		return err
	}
	if row == 0 {
		return errors.New("can not update account by this user name:" + a.UserName)
	}
	return nil
}

func UpdateSignOn(a *domain.Account) error {
	d, err := util.GetConnection()
	if err != nil {
		return err
	}
	r, err := d.Exec(updateSigOnSQL, a.Password, a.UserName)
	if err != nil {
		return err
	}
	row, err := r.RowsAffected()
	if err != nil {
		return err
	}
	if row == 0 {
		return errors.New("can not update account by this user name:" + a.UserName)
	}
	return nil
}
