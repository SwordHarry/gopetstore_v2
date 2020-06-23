package controller

import (
	"github.com/gin-gonic/gin"
	"gopetstore_v2/src/config"
	"gopetstore_v2/src/domain"
	"gopetstore_v2/src/service"
	"gopetstore_v2/src/util"
	"log"
	"net/http"
)

// file name
const (
	signInFormFile      = "signInForm.html"
	registerFormFile    = "registerForm.html"
	editAccountFormFile = "editAccountForm.html"
	mainFile            = "main.html"
)

// view select value
var (
	languages = []string{
		"english",
		"japanese",
	}
	categories = []string{
		"FISH",
		"DOGS",
		"REPTILES",
		"CATS",
		"BIRDS",
	}
)

// view
// view login form
func ViewLogin(c *gin.Context) {
	c.HTML(http.StatusOK, signInFormFile, gin.H{})
}

// view register form
func ViewRegister(c *gin.Context) {
	c.HTML(http.StatusOK, registerFormFile, gin.H{
		"Languages":  languages,
		"Categories": categories,
	})
}

// view edit account form
func ViewEditAccount(c *gin.Context) {
	a := util.GetAccountFromSession(c.Request)
	c.HTML(http.StatusOK, editAccountFormFile, gin.H{
		"Account":    a,
		"Languages":  languages,
		"Categories": categories,
	})
}

// action
// login
func Login(c *gin.Context) {
	userName := c.PostForm("username")
	password := c.PostForm("password")
	a, err := service.GetAccountByUserNameAndPassword(userName, password)
	if err != nil {
		util.ViewError(c, err)
		return
	}
	if a != nil {
		s, err := util.GetSession(c.Request)
		if err != nil {
			util.ViewError(c, err)
			return
		}
		if s != nil {
			err = s.Save(config.AccountKey, a, c.Writer, c.Request)
			if err != nil {
				util.ViewError(c, err)
				return
			}
		}
		c.HTML(http.StatusOK, mainFile, gin.H{
			"Account": a,
		})
	} else {
		c.HTML(http.StatusOK, signInFormFile, gin.H{
			"Message": "登录失败，账号或密码错误",
		})
	}
}

// sign out
func SignOut(c *gin.Context) {
	s, err := util.GetSession(c.Request)
	if err != nil {
		util.ViewError(c, err)
		return
	}
	if s != nil {
		err = s.Del(config.AccountKey, c.Writer, c.Request)
		err = s.Del(config.CartKey, c.Writer, c.Request)
		err = s.Del(config.OrderKey, c.Writer, c.Request)
		if err != nil {
			util.ViewError(c, err)
			return
		}
	}
	c.HTML(http.StatusOK, mainFile, gin.H{})
}

// register
func NewAccount(c *gin.Context) {
	accountInfo := getAccountFromInfoForm(c)
	repeatedPassword := c.PostForm("repeatedPassword")
	if accountInfo.Password != repeatedPassword {
		c.HTML(http.StatusOK, registerFormFile, gin.H{
			"Message":    "密码和重复密码不一致",
			"Languages":  languages,
			"Categories": categories,
		})
		return
	}
	a, err := service.GetAccountByUserName(accountInfo.UserName)
	if err != nil {
		util.ViewError(c, err)
		return
	}
	if a == nil {
		// 进行注册
		err := service.InsertAccount(accountInfo)
		if err != nil {
			util.ViewError(c, err)
		} else {
			c.HTML(http.StatusOK, signInFormFile, gin.H{
				"Message": "注册成功",
			})
		}
	} else {
		c.HTML(http.StatusOK, registerFormFile, gin.H{
			"Message":    "该用户名已存在",
			"Languages":  languages,
			"Categories": categories,
		})
	}
}

// update account
func ConfirmEdit(c *gin.Context) {
	a := getAccountFromInfoForm(c)
	err := service.UpdateAccount(a)
	if err != nil {
		util.ViewError(c, err)
		return
	}
	// 修改成功后需要重置 session
	s, err := util.GetSession(c.Request)
	if err != nil {
		log.Printf("ConfirmEdit GetSession error: %v", err.Error())
	}
	if s != nil {
		err = s.Save(config.AccountKey, a, c.Writer, c.Request)
		if err != nil {
			log.Printf("ConfirmEdit Save error: %v", err.Error())
		}
	}
	c.HTML(http.StatusOK, editAccountFormFile, gin.H{
		"Message": "修改成功",
		"Account": a,
	})
}

// get account info from form
func getAccountFromInfoForm(c *gin.Context) *domain.Account {
	userName := c.PostForm("username")
	password := c.PostForm("password")
	firstName := c.PostForm("firstName")
	lastName := c.PostForm("lastName")
	email := c.PostForm("email")
	phone := c.PostForm("phone")
	address1 := c.PostForm("address1")
	address2 := c.PostForm("address2")
	city := c.PostForm("city")
	state := c.PostForm("state")
	zip := c.PostForm("zip")
	country := c.PostForm("country")
	languagePreference := c.PostForm("languagePreference")
	favouriteCategoryId := c.PostForm("favouriteCategoryId")
	listOption := c.PostForm("listOption")
	bannerOption := c.PostForm("bannerOption")

	finalListOption := len(listOption) > 0
	finalBannerOption := len(bannerOption) > 0
	a := &domain.Account{
		UserName:            userName,
		Email:               email,
		FirstName:           firstName,
		LastName:            lastName,
		Status:              "OK",
		Address1:            address1,
		Address2:            address2,
		City:                city,
		State:               state,
		Zip:                 zip,
		Country:             country,
		Phone:               phone,
		Password:            password,
		FavouriteCategoryId: favouriteCategoryId,
		LanguagePreference:  languagePreference,
		ListOption:          finalListOption,
		BannerOption:        finalBannerOption,
	}
	return a
}
