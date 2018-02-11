package ykredis

import (
	"fmt"
	"net"

	cons "../ykconstant"
	redigo "github.com/garyburd/redigo/redis"
)

type User struct {
	Uid        cons.UID
	Nick       string
	LastLogin  string
	FirstLogin string
}

/*
	在线列表
*/
type IUserOL interface {
	Create()
	SetOL(conn *net.Conn, uid cons.UID)
	DelOL(uid cons.UID)
}

type UserOL struct {
	OLInfo map[cons.UID](*net.Conn)
}

func (yk *UserOL) Create() {
	yk.OLInfo = make(map[cons.UID](*net.Conn), cons.MAX_OL)
}

func (yk *UserOL) SetOL(conn *net.Conn, uid cons.UID) {
	fmt.Println("set uid=", uid)
	yk.OLInfo[uid] = conn
}

func (yk *UserOL) DelOL(uid cons.UID) {
	fmt.Println("del uid=", uid)
	yk.OLInfo[uid] = nil
}

func (yk *UserOL) GetOL(uid cons.UID) *net.Conn {
	return yk.OLInfo[uid]
}

/*
redis存储user
*/

var (
	redCfg = Config{"127.0.0.1:5276", "", 2, 1, 0, false, "yk", ":", "?"}
)

type UserCenterStorge struct {
	red *redigo.Conn
}

func (yk *UserCenterStorge) Get() {
	pool := NewRConnectionPool(redCfg)
	red := pool.Get()
	yk.red = &red
}

func (yk *UserCenterStorge) GetUser(uid cons.UID) (user User) {
	return user
}

func (yk *UserCenterStorge) SetUser(user User) {

}
