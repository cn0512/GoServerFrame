package Svrs

/*
	1,服务注册
*/

import (
	tl "MSvrs/Core/Utils"
	"encoding/json"
	"fmt"
)

type IReg interface {
	Reg()
	Unreg()
	Encode() []byte
	Decode()
}

type SvrItem struct {
	SvrType string
	Uuid    string
	Addr    string //ip:port
	State   bool   //mark:0,disbale;1,enable
}

var MSvrsItems map[string]SvrItem

func init() {
	MSvrsItems = make(map[string]SvrItem, 8)
}

func (svr *SvrItem) Reg() {
	id := tl.UU()
	svr.Uuid = id
	svr.State = true
	MSvrsItems[id] = *svr
	fmt.Println(svr, "is reg")
}

func (svr *SvrItem) Unreg() {
	v, ok := MSvrsItems[svr.Uuid]
	if !ok {
		fmt.Println(svr, "err unreg")
		return
	}
	svr.State = false
	v.State = false
	fmt.Println(svr, v, "is unreg")
}

func (svr *SvrItem) Encode() []byte {
	//json
	buf, err := json.Marshal(*svr)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return buf
}

func (svr *SvrItem) Decode(buf []byte)error  {
	err := json.Unmarshal(buf, svr)
	if err != nil {
		tl.Logout("%v", err)
	}
	return err
}
