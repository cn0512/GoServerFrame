package main

/*
	1, 服务注册,发现服务
	2, redis/pub.sub
*/

import (
	"fmt"

	"github.com/cn0512/GoServerFrame/Config"
	redis "github.com/cn0512/GoServerFrame/Core/DB/Redis"
	mq "github.com/cn0512/GoServerFrame/Core/MQ/Redis"
	svrCfg "github.com/cn0512/GoServerFrame/Core/Svrs"
	"github.com/cn0512/GoServerFrame/Core/Utils"

	"os"
	"os/signal"
)

type Msg struct {
}

func (m *Msg) Call(msg string) {
	//fmt.Println("Call:", msg)
	//fmt.Println("\n")
	item := svrCfg.SvrItem{}
	//fmt.Println(item)
	err := item.Decode([]byte(msg))
	if err != nil {
		return
	}
	//fmt.Println(item)
	if item.State == true {
		item.Reg()
	} else {
		item.Unreg()
	}

	go Save2Redis(item)
}

var redis_cfg = redis.Config{
	Server:    "localhost:6379",
	MaxIdle:   1,
	MaxActive: 0,
}

func Save2Redis(item svrCfg.SvrItem) {
	pool := redis.NewRConnectionPool(redis_cfg)
	con := pool.Get()
	if item.State {
		redis.HSet(&con, "reg:server", item.Uuid, item)
	} else {
		redis.HDel(&con, "reg:server", item.Uuid)
	}

}

func main() {
	fmt.Println("reg_svr is running")
	//【1】sub
	Utils.Logout("try Sub`s redis")
	sub_reg, err1 := mq.NewSub(Config.Topic_Svrs_Reg, &Msg{})
	sub_unreg, err2 := mq.NewSub(Config.Topic_Svrs_Unreg, &Msg{})
	defer sub_reg.Shutdown()
	defer sub_unreg.Shutdown()
	if err1 != nil || err2 != nil {
		fmt.Println(err1, err2)
		return
	}
	Utils.Logout("finsh Sub`s redis")

	//【2】mq

	qt := make(chan os.Signal, 1)
	signal.Notify(qt, os.Interrupt, os.Kill)
	<-qt
}
