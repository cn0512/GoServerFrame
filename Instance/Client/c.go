package main

import (
	RPC "MSvrs/Core/RPC"
	"fmt"
	"log"
	"net/rpc/jsonrpc"

	cfg "MSvrs/Config"
)

func main() {
	//连接远程rpc服务
	rpc, err := jsonrpc.Dial("tcp", cfg.Json_rpc_addr)
	if err != nil {
		log.Fatal(err)
	}
	ret := RPC.CreateUserRep{}
	//调用远程方法
	user := RPC.CreateUserReq{
		Uuid:    "",
		Account: "yk1",
		Nick:    "yk1",
		Passwd:  "",
	}

	err2 := rpc.Call("Rpc_logic.CreateUser", user, &ret)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println(ret)
}
