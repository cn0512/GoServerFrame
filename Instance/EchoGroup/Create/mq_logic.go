package main

import (
	RPC "MSvrs/Core/RPC"
	"fmt"
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
