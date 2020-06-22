package util

import (
	"github.com/gorilla/sessions"
	"gopetstore_v2/src/config"
	"gopetstore_v2/src/domain"
	"log"
	"net/http"
)

/*
对 sessions 库的再封装,实现简单session功能
*/
type Session struct {
	se *sessions.Session
}

// 秘钥，生成唯一 sessionStore
const secretKey = "go-pet-store"

// go web 标准库没有 session，需要自己开发封装或使用第三方的库
var sessionStore = sessions.NewFilesystemStore("", []byte(secretKey))

const sessionName = "session"

func init() {
	// 设置 fileSystemStore 的最大存储长度，防止溢出
	sessionStore.MaxLength(5 * 4096)
}

// 初始化，通过这个获取唯一 session
func GetSession(r *http.Request) (*Session, error) {

	s, err := sessionStore.Get(r, sessionName)
	if err != nil {
		return nil, err
	}
	return &Session{
		s,
	}, nil
}

// 存储和更新，复杂类型存储前需要 gob.Register 进行序列化
func (s *Session) Save(key string, val interface{}, w http.ResponseWriter, r *http.Request) error {
	s.se.Values[key] = val
	return s.se.Save(r, w)
}

// 获取值
func (s *Session) Get(key string) (result interface{}, ok bool) {
	result, ok = s.se.Values[key]
	return
}

// 删除值
func (s *Session) Del(key string, w http.ResponseWriter, r *http.Request) error {
	delete(s.se.Values, key)
	return s.se.Save(r, w)
}

// 从 session 中获取 account
func GetAccountFromSession(r *http.Request) *domain.Account {
	s, err := GetSession(r)
	if err != nil {
		log.Printf("get session error: %v", err.Error())
		return nil
	}
	if s != nil {
		r, ok := s.Get(config.AccountKey)
		if !ok {
			// account 不存在，已登出
			return nil
		}
		a, ok := r.(*domain.Account)
		if !ok {
			log.Print("type assert error *domain.Account")
			return nil
		}
		return a
	}
	log.Print("session get account error: session is nil")
	return nil
}

// 从session中获取cart
func GetCartFromSessionAndSave(w http.ResponseWriter, r *http.Request, callback func(cart *domain.Cart)) *domain.Cart {
	// 使用 session 存储 cart 购物车
	s, err := GetSession(r)
	if err != nil {
		log.Printf("session error for getSession: %v", err.Error())
	}
	var cart *domain.Cart
	// 成功生成 session
	if s != nil {
		c, ok := s.Get(config.CartKey)
		if !ok {
			// 初始化 购物车
			c = domain.NewCart()
		}
		// 调用回调对cart 进行操作
		cart, ok = c.(*domain.Cart)
		if ok && callback != nil {
			callback(cart)
		}
		// 将新的购物车进行存储覆盖
		err := s.Save(config.CartKey, c, w, r)
		if err != nil {
			log.Printf("GetCartFromSessionAndSave session error for Save: %v", err.Error())
		}
	}
	return cart
}

// 从 session 中获取 order
func GetOrderFromSessionAndSave(w http.ResponseWriter, r *http.Request, callback func(order *domain.Order)) *domain.Order {
	// 使用 session 存储 order 订单
	s, err := GetSession(r)
	if err != nil {
		log.Printf("session error for getSession: %v", err.Error())
	}
	var order *domain.Order
	// 成功生成 session
	if s != nil {
		c, ok := s.Get(config.OrderKey)
		if !ok {
			return nil
		}
		// 调用回调对 order 进行操作
		order, ok = c.(*domain.Order)
		if ok && callback != nil {
			callback(order)
		}
		// 将新的购物车进行存储覆盖
		err := s.Save(config.OrderKey, c, w, r)
		if err != nil {
			log.Printf("GetCartFromSessionAndSave session error for Save: %v", err.Error())
		}
	}
	return order
}
