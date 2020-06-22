package persistence

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"gopetstore_v2/src/domain"
	"gopetstore_v2/src/util"
)

const getAccountByUsernameSQL = `SELECT SIGNON.USERNAME as userid,ACCOUNT.EMAIL as email,ACCOUNT.FIRSTNAME as firstname,
ACCOUNT.LASTNAME as lastname,ACCOUNT.STATUS as status,ACCOUNT.ADDR1 AS addr1,ACCOUNT.ADDR2 AS addr2,ACCOUNT.CITY as city,
ACCOUNT.STATE as state,ACCOUNT.ZIP as zip,ACCOUNT.COUNTRY as country,ACCOUNT.PHONE as phone,PROFILE.LANGPREF AS langpref,
PROFILE.FAVCATEGORY AS favcategory,PROFILE.MYLISTOPT AS mylistopt,PROFILE.BANNEROPT AS banneropt,
BANNERDATA.BANNERNAME as bannername FROM ACCOUNT, PROFILE, SIGNON, BANNERDATA
WHERE ACCOUNT.USERID = ? AND SIGNON.USERNAME = ACCOUNT.USERID AND PROFILE.USERID = ACCOUNT.USERID AND
 PROFILE.FAVCATEGORY = BANNERDATA.FAVCATEGORY`

// get account by userName and password from signOn, account, bannerData
const getAccountByUsernameAndPasswordSQL = `SELECT SIGNON.USERNAME as userid,ACCOUNT.EMAIL as email,
ACCOUNT.FIRSTNAME as firstname,ACCOUNT.LASTNAME as lastname,ACCOUNT.STATUS as status,ACCOUNT.ADDR1 AS addr1,
ACCOUNT.ADDR2 AS addr2,ACCOUNT.CITY as city,ACCOUNT.STATE as state,ACCOUNT.ZIP as zip,ACCOUNT.COUNTRY as country,
ACCOUNT.PHONE as phone,PROFILE.LANGPREF AS langpref,PROFILE.FAVCATEGORY AS favcategory,PROFILE.MYLISTOPT AS mylistopt,
PROFILE.BANNEROPT AS banneropt,BANNERDATA.BANNERNAME as bannername FROM ACCOUNT, PROFILE, SIGNON, BANNERDATA 
WHERE ACCOUNT.USERID = ? AND SIGNON.PASSWORD = ? AND SIGNON.USERNAME = ACCOUNT.USERID AND 
PROFILE.USERID = ACCOUNT.USERID AND PROFILE.FAVCATEGORY = BANNERDATA.FAVCATEGORY`

// Insert
// insert account from account
const insertAccountSQL = `INSERT INTO ACCOUNT (EMAIL, FIRSTNAME, LASTNAME, STATUS, ADDR1, ADDR2, CITY, STATE, ZIP, COUNTRY, PHONE, USERID) 
VALUES(:email, :firstname, :lastname, :status, :addr1, :addr2, :city, :state, :zip, :country, :phone, :userid)`

// insert profile from profile
const insertProfileSQL = `INSERT INTO PROFILE (LANGPREF, FAVCATEGORY, USERID, mylistopt, banneropt) 
VALUES (:langpref, :favcategory, :userid, :mylistopt, :banneropt)`

// insert username and password from signOn
const insertSigOnSQL = `INSERT INTO SIGNON (USERNAME,PASSWORD) VALUES (:userid, :password)`

// Update
// update account from account
const updateAccountSQL = `UPDATE ACCOUNT SET EMAIL = :email,FIRSTNAME = :firstname,LASTNAME = :lastname,
STATUS = :status,ADDR1 = :addr1,ADDR2 = :addr2,CITY = :city,STATE = :state,ZIP = :zip,
COUNTRY = :country,PHONE = :phone WHERE USERID = :userid`

// update profile from profile
const updateProfileSQL = `UPDATE PROFILE SET LANGPREF = :langpref, FAVCATEGORY = :favcategory,mylistopt = :mylistopt,
banneropt = :banneropt WHERE USERID = :userid`

// update password by userName from signOn
const updateSigOnSQL = `UPDATE SIGNON SET PASSWORD = :password WHERE USERNAME = :userid`

// query
// get account by user name
// 找不到则意味着可以注册，而不是报错
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
		if err == sql.ErrNoRows {
			return nil, nil
		}
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
		// 账号或密码错误
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return a, nil
}

// insert
// 插入 account 数据表
func InsertAccount(a *domain.Account) error {
	// 使用事务，插入时如果失败则回滚
	return util.ExecTransaction(func(tx *sqlx.Tx) error {
		_, err := tx.NamedExec(insertAccountSQL, a)
		if err != nil {
			tx.Rollback()
			return err
		}
		_, err = tx.NamedExec(insertProfileSQL, a)
		if err != nil {
			tx.Rollback()
			return err
		}
		_, err = tx.NamedExec(insertSigOnSQL, a)
		if err != nil {
			tx.Rollback()
			return err
		}
		return nil
	})

}

//
//// 插入 profile 数据表
//func InsertProfile(a *domain.Account) error {
//	d, err := util.GetConnection()
//	defer func() {
//		_ = d.Close()
//	}()
//	if err != nil {
//		return err
//	}
//	r, err := d.NamedExec(insertProfileSQL, a)
//	if err != nil {
//		return err
//	}
//	row, err := r.RowsAffected()
//	if err != nil {
//		return err
//	}
//	if row == 0 {
//		return errors.New("can not insert profile by this user name:" + a.UserName)
//	}
//	return nil
//}
//
//// 插入 signon 数据表
//func InsertSignOn(a *domain.Account) error {
//	d, err := util.GetConnection()
//	defer func() {
//		_ = d.Close()
//	}()
//	if err != nil {
//		return err
//	}
//	r, err := d.NamedExec(insertSigOnSQL, a)
//	if err != nil {
//		return err
//	}
//	row, err := r.RowsAffected()
//	if err != nil {
//		return err
//	}
//	if row == 0 {
//		return errors.New("can not insert sign on by this user name:" + a.UserName)
//	}
//	return nil
//}

// update
func UpdateAccount(a *domain.Account) error {
	return util.ExecTransaction(func(tx *sqlx.Tx) error {
		_, err := tx.NamedExec(updateAccountSQL, a)
		if err != nil {
			tx.Rollback()
			return err
		}
		_, err = tx.NamedExec(updateProfileSQL, a)
		if err != nil {
			tx.Rollback()
			return err
		}
		_, err = tx.NamedExec(updateSigOnSQL, a)
		if err != nil {
			tx.Rollback()
			return err
		}
		return nil
	})
}

//
//func UpdateProfile(a *domain.Account) error {
//	d, err := util.GetConnection()
//	if err != nil {
//		return err
//	}
//	r, err := d.NamedExec(updateProfileSQL, a)
//	if err != nil {
//		return err
//	}
//	_, err = r.RowsAffected()
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//func UpdateSignOn(a *domain.Account) error {
//	d, err := util.GetConnection()
//	if err != nil {
//		return err
//	}
//	r, err := d.NamedExec(updateSigOnSQL, a)
//	if err != nil {
//		return err
//	}
//	_, err = r.RowsAffected()
//	if err != nil {
//		return err
//	}
//	return nil
//}
