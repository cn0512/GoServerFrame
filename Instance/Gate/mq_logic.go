package main

import (
	"fmt"

	ykmq "github.com/cn0512/GoServerFrame/Core/MQ/YKMQ"
	RPC "github.com/cn0512/GoServerFrame/Core/RPC"
)

type MsgCreateUserRep struct {
}

func (m *MsgCreateUserRep) Call(msg string) {
	fmt.Println("Call:", msg)

	u := RPC.CreateUserRep{}
	err := u.Decode([]byte(msg))
	if err != nil {
		panic(err)
	}
	//插入队列
	ykmq.Put(u)
}
