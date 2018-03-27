package main

/*
	1,use redis-mq sub/pub,
	2,create user then return create-result
*/

import (
	mq "github.com/cn0512/GoServerFrame/Core/MQ/Redis"
	RPC "github.com/cn0512/GoServerFrame/Core/RPC"

	ps "github.com/aalness/go-redis-pubsub"

	"os"
	"os/signal"

	"github.com/cn0512/GoServerFrame/Config"
	"github.com/cn0512/GoServerFrame/Core/Utils"
)

var pub ps.Publisher

func Pubresult(u RPC.CreateUserRep) {
	pub.Publish(Config.Topic_Svrs_CreateUser_Rep, u.Encode())
}

func main() {
	//【1】pub mq
	pub = mq.NewPub()
	defer pub.Shutdown()
	Utils.Logout("Pub`s init")
	//sub mq
	sub, err_sub := mq.NewSub(Config.Topic_Svrs_CreateUser_Req, &MsgCreateUserReq{})
	if err_sub != nil {
		panic(err_sub)
		return
	}
	defer sub.Shutdown()

	qt := make(chan os.Signal, 1)
	signal.Notify(qt, os.Interrupt, os.Kill)
	<-qt
}
