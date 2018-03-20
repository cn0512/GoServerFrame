package main

import (
	RPC "MSvrs/Core/RPC"
	"fmt"
	ykmq "MSvrs/Core/MQ/YKMQ"
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
