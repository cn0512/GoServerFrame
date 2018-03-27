package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"

	RPC "github.com/cn0512/GoServerFrame/Core/RPC"

	cfg "github.com/cn0512/GoServerFrame/Config"
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
