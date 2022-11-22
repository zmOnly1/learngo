package model

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

var (
	UserTable = "users"
)

type UserMgr struct {
	pool *redis.Pool
}

func NewUserMgr(pool *redis.Pool) *UserMgr {
	mgr := &UserMgr{
		pool: pool,
	}
	return mgr
}

func (p *UserMgr) getUser(conn redis.Conn, id int) (*User, error) {
	result, err := redis.String(conn.Do("HGet", UserTable, fmt.Sprintf("%d", id)))
	if err != nil {
		if err == redis.ErrNil {
			return nil, ErrUserNotExist
		}
		return nil, err
	}
	user := &User{}
	err = json.Unmarshal([]byte(result), user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (p *UserMgr) Login(id int, passwd string) (*User, error) {
	conn := p.pool.Get()
	defer conn.Close()

	user, err := p.getUser(conn, id)
	if err != nil {
		return nil, err
	}
	if user.UserId != id || user.Passwd != passwd {
		return nil, ErrInvalidPasswd
	}
	user.Status = UserStatusOnline
	user.LastLogin = fmt.Sprintf("%v", time.Now())
	return user, nil
}

func (p *UserMgr) Register(user *User) error {
	conn := p.pool.Get()
	defer conn.Close()
	if user == nil {
		fmt.Println("invalid user")
		return ErrInvalidParams
	}
	_, err := p.getUser(conn, user.UserId)
	if err != nil {
		return ErrUserExist
	}
	if err != ErrUserNotExist {
		return err
	}
	data, err := json.Marshal(user)
	if err != nil {
		return err
	}
	_, err = conn.Do("HSet", UserTable, fmt.Sprintf("%d", user.UserId), string(data))
	if err != nil {
		return err
	}
	return nil

}
