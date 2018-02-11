package ykconstant

import (
	"bytes"
	"encoding/binary"
)

type UID int32
type Entity int32

const (
	Version            = "0.0.1"
	TCP_IP             = "127.0.0.1:6018"
	Buf_Size           = 4096
	Secret_key         = "ykgame" //异或密钥
	MAX_OL             = 0x0fffff
	AMQP_uri           = "amqp://guest:guest@localhost:5672"
	AMQP_exchangeName  = "exc_c2s"
	AMQP_exchangeName2 = "exc_s2c"
	AMQP_exchangeType  = "fanout"
	AMQP_exchangeType2 = "direct"
	AMQP_routingKey    = "c2s_key"
	AMQP_queueName     = "c2s_queue"
	AMQP_routingKey2   = "s2c_key"
	AMQP_queueName2    = "s2c_queue"
	AMQP_reliable      = true
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

//角色
const (
	Entity_None = iota
	Entity_Client
	Entity_Web
	Entity_Any   = 10
	Entity_Login = 100
	Entity_Game  = 200
	Entity_DB    = 300
	Entity_Lobby = 400
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
