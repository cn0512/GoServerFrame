package RPC

import (
	"encoding/json"
	"fmt"
	"strings"

	tl "github.com/cn0512/GoServerFrame/Core/Utils"
)

type User struct {
	Id       int64
	Account  string
	Passwd   string
	Nickname string
	RegTime  string
}

type UserDetail struct {
	id    int64
	age   int
	phone string
	email string
}

func (user User) Equal(o User) bool {
	if user.Id == o.Id && o.Id != 0 {
		return true
	}

	if strings.EqualFold(user.Account, o.Account) && !tl.IsBlank(o.Account) {
		return true
	}

	return false
}

func (user *User) Encode() []byte {
	//json
	buf, err := json.Marshal(*user)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return buf
}

func (user *User) Decode(buf []byte) error {
	err := json.Unmarshal(buf, user)
	if err != nil {
		tl.Logout("%v", err)
	}
	return err
}
