package route

// 和 session 有关的中间件，负责将数据在 session 和 context 之间进行流动

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gopetstore_v2/src/domain"
	"gopetstore_v2/src/util"
	"log"
)

// session 中间件
func SessionMiddleWare(key ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		s, err := util.GetSession(c.Request)
		if err != nil {
			log.Printf("get session error: %v", err.Error())
		}
		if s != nil {
			for _, k := range key {
				setValToContextFromSession(k, c, s)
			}
		}
		// 先将 session 中的数据放到 context 中
		c.Next()
		// 再将 context 中的数据存到 session 中
		if s != nil {
			for _, k := range key {
				err = setValToSessionFromContext(k, c, s)
				if err != nil {
					log.Printf("setValToSessionFromContext error: %v", err.Error())
				}
			}
		}
	}
}

// 是否登录，从 session 中获取 并往 context 中存放 account 指针
func AccountLogin(c *gin.Context) {
	const account = "account"
	s, err := util.GetSession(c.Request)
	if err != nil {
		log.Printf("get session error: %v", err.Error())
	}
	if s != nil {
		a, ok := s.Get(account)
		if ok {
			c.Set(account, a)
		} else {
			c.Set(account, nil)
		}
	}
	c.Next()
}

// 从 session 中取值，并存入到 context 中
func setValToContextFromSession(key string, c *gin.Context, s *util.Session) {
	v, ok := s.Get(key)
	if ok {
		switch v.(type) {
		case *domain.Account:
			c.Set(key, v)
		case *domain.Product:
			c.Set(key, v)
		case *domain.Cart:
			c.Set(key, v)
		case *domain.Order:
			c.Set(key, v)
		default:
			log.Printf("type v is %T:", v)
		}
	}
}

// 从 context 中取值，并存入到 session 中
func setValToSessionFromContext(key string, c *gin.Context, s *util.Session) error {
	v, ok := c.Get(key)
	if ok {
		switch v.(type) {
		case *domain.Account:
			return s.Save(key, v, c.Writer, c.Request)
		case *domain.Product:
			return s.Save(key, v, c.Writer, c.Request)
		case *domain.Cart:
			return s.Save(key, v, c.Writer, c.Request)
		case *domain.Order:
			return s.Save(key, v, c.Writer, c.Request)
		default:
			log.Printf("type v is %T:", v)
			return errors.New("type v is out of the domain")
		}
	}
	return nil
}
