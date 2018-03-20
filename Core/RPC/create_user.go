package RPC

import (
	tl "MSvrs/Core/Utils"
	"encoding/json"
	"fmt"
)

type CreateUserReq struct {
	Uuid    string
	Account string
	Passwd  string
	Nick    string
}

type CreateUserRep struct {
	Uuid string
	Ret  int
	ID   int64
}

func (user *CreateUserReq) Encode() []byte {
	//json
	buf, err := json.Marshal(*user)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return buf
}

func (user *CreateUserReq) Decode(buf []byte) error {
	err := json.Unmarshal(buf, user)
	if err != nil {
		tl.Logout("%v", err)
	}
	return err
}

func (user *CreateUserRep) Encode() []byte {
	//json
	buf, err := json.Marshal(*user)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return buf
}

func (user *CreateUserRep) Decode(buf []byte) error {
	err := json.Unmarshal(buf, user)
	if err != nil {
		tl.Logout("%v", err)
	}
	return err
}
