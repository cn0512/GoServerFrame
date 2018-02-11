package main

import (
	"flag"
	"fmt"

	c "./ykconstant"
	//net "./yknet"
	"github.com/golang/protobuf/proto"

	ykmsg "./ykmsg"
	"./ykprotoco"
)

var (
	version = flag.String("v", "0.0.1", "version = ?")
	tcp_ip  = flag.String("ip", "127.0.0.1:6018", "port = ?")
)

type gobSt struct {
	Age  int
	Name string
}

func main() {
	flag.Parse()
	fmt.Println(c.Version)
	fmt.Println(c.TCP_IP)
	fmt.Println(*version)
	fmt.Println(*tcp_ip)

	header := ykmsg.MsgHeader{1029, 20, 1}
	headerbuf := header.Get()
	fmt.Println(headerbuf)

	var login YKGameMsg.LoginMsgReq
	login.Uid = 1000
	login.Pwd = "passord"
	login.CheckCode = "code"
	buf := login.String()
	fmt.Println("login1=", login, buf)
	data, _ := proto.Marshal(&login)
	fmt.Println("login1,buf=", data)
	var login2 YKGameMsg.LoginMsgReq
	proto.Unmarshal(data, &login2)
	fmt.Println("login2=", login2)

	var begin YKGameMsg.GameBegin
	begin.Seat = 1
	begin.Players = make([]*YKGameMsg.StPlayer, 2)
	cards0 := []int32{1, 2, 3, 4, 5}
	begin.Players[0] = &YKGameMsg.StPlayer{1, cards0}

	fmt.Println(begin)

	/*
		Gob
	*/
	//encode
	st := gobSt{12, "123"}
	var ykgob ykmsg.GobMsg
	bufMsg, err := ykgob.Encode(&st, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("encode=", bufMsg)
	//decode
	var st2 gobSt
	err = ykgob.Decode(bufMsg, &st2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("decode=", st2)
}
