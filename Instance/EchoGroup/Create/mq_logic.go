package main

import (
	"fmt"

	RPC "github.com/cn0512/GoServerFrame/Core/RPC"
)

type MsgCreateUserReq struct {
}

func (m *MsgCreateUserReq) Call(msg string) {
	fmt.Println("Call:", msg)

	u := RPC.CreateUserReq{}
	err := u.Decode([]byte(msg))
	if err != nil {
		panic(err)
	}

	//check the db
	ret := CheckDatabase(u)

	//pub the result
	Pubresult(ret)
}
