package YKMQ

import (
	"sync"

	RPC "github.com/cn0512/GoServerFrame/Core/RPC"
)

var safemap sync.Map

func init() {

}

func Reset(uuid string) {
	safemap.Store(uuid, RPC.CreateUserRep{Uuid: uuid, Ret: 400, ID: 0})
}

func Get(uuid string) (RPC.CreateUserRep, bool) {
	if v, ok := safemap.Load(uuid); ok && v.(interface{}).(RPC.CreateUserRep).Ret != 400 {
		s := v.(interface{}).(RPC.CreateUserRep)
		safemap.Delete(uuid)
		return s, true
	}
	return RPC.CreateUserRep{}, false
}

func Put(s RPC.CreateUserRep) {
	safemap.Store(s.Uuid, s)
}
