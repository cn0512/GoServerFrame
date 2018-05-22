package Yklua

import (
	"github.com/yuin/gopher-lua"
)

func init() {

}

func Load(file string) *lua.LState {
	L := lua.NewState()
	defer L.Close()
	if err := L.DoFile(file); err != nil {
		panic(err)
	}
	return L
}
