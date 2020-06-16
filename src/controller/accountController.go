package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// view login form
func ViewLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "signInForm.html", gin.H{})
}

// view register form
func ViewRegister(c *gin.Context) {
	c.HTML(http.StatusOK, "registerForm.html", gin.H{})
}
