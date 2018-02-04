package ykconstant

import (
	"bytes"
	"encoding/binary"
)

const (
	Version           = "0.0.1"
	TCP_IP            = "127.0.0.1:6018"
	Buf_Size          = 4096
	Secret_key        = "ykgame" //异或密钥
	AMQP_uri          = "amqp://guest:guest@localhost:5672"
	AMQP_exchangeName = "ykgame"
	AMQP_exchangeType = "direct"
	AMQP_routingKey   = "ykgame_key"
	AMQP_reliable     = true
)

//c/s 交互协议ID
const (
	CS_MSGID_LOGIN = iota
	CS_MSGID_LOGOUT
	CS_MSGID_PLAY
	CS_MSGID_GAMEBEGIN
	CS_MSGID_GAMEEND
)

//router 协议ID
const (
	R_MSGID_LOGIN = 1000 + iota
	R_MSGID_GAMEBEGIN
	R_MSGID_GAMEEND
)

//服务器角色
const (
	Role_Login = iota
	Role_Game
	Role_DB
	Role_Lobby
)

//整形转换成字节
func IntToBytes(n int32) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, n)
	return bytesBuffer.Bytes()
}

//字节转换成整形
func BytesToInt(b []byte) int32 {
	bytesBuffer := bytes.NewBuffer(b)
	var tmp int32
	binary.Read(bytesBuffer, binary.BigEndian, &tmp)
	return tmp
}

type Role int

//router 核心
type MsgRouter struct {
	src Role
	dst Role
}
