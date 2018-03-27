package main

/*
	1. S- S :use redis:pub/sub trans msg between svrs
	2, C - S: JsonRPC
*/

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
	"os/signal"

	ps "github.com/aalness/go-redis-pubsub"
	"github.com/cn0512/GoServerFrame/Config"
	mq "github.com/cn0512/GoServerFrame/Core/MQ/Redis"
	"github.com/cn0512/GoServerFrame/Core/Utils"
)

func chkError(err error) {
	if err != nil {
		Utils.Logout("%v", err)
	}
}

var pub ps.Publisher

func init() {

}

func main() {
	fmt.Println("gate is running")

	//【1】pub mq
	pub = mq.NewPub()
	defer pub.Shutdown()
	Utils.Logout("Pub`s init")
	//sub mq
	sub, err_sub := mq.NewSub(Config.Topic_Svrs_CreateUser_Rep, &MsgCreateUserRep{})
	if err_sub != nil {
		panic(err_sub)
		return
	}
	defer sub.Shutdown()

	//【2】rpc
	logic := new(Rpc_logic)
	//注册rpc服务
	rpc.Register(logic)
	//获取tcpaddr
	tcpaddr, err := net.ResolveTCPAddr("tcp4", Config.Json_rpc_addr)
	chkError(err)
	//监听端口
	tcplisten, err2 := net.ListenTCP("tcp", tcpaddr)
	chkError(err2)
	for {
		conn, err3 := tcplisten.Accept()
		if err3 != nil {
			continue
		}
		//使用goroutine单独处理rpc连接请求
		//这里使用jsonrpc进行处理
		go jsonrpc.ServeConn(conn)
	}

	qt := make(chan os.Signal, 1)
	signal.Notify(qt, os.Interrupt, os.Kill)
	<-qt
}
