package main

import (
	cfg "github.com/cn0512/GoServerFrame/Config"
	//mysqlx "github.com/cn0512/GoServerFrame/Core/DB/Mysqlx"
	"context"
	"fmt"

	RPC "github.com/cn0512/GoServerFrame/Core/RPC"
	//"github.com/jmoiron/sqlx"
	"time"

	ykmq "github.com/cn0512/GoServerFrame/Core/MQ/YKMQ"
	tl "github.com/cn0512/GoServerFrame/Core/Utils"
)

type Rpc_logic struct {
	ctx context.Context
}

//var DB *sqlx.DB = mysqlx.New()

func (lg *Rpc_logic) CreateUser(user RPC.CreateUserReq, ret *RPC.CreateUserRep) error {

	fmt.Println(user)

	//[1]pub create msg
	user.Uuid = "CreateUser:" + tl.UU()
	pub.Publish(cfg.Topic_Svrs_CreateUser_Req, user.Encode())
	lg.ctx = context.TODO()
	//[2]Sync
	ykmq.Reset(user.Uuid)
	var ok bool
	var rep RPC.CreateUserRep
	for {
		if rep, ok = ykmq.Get(user.Uuid); !ok {
			time.Sleep(300 * time.Microsecond)
			//time.Sleep(1 * time.Second)
			//fmt.Println("sleep", time.Now())
		} else {
			break
		}
	}
	*ret = rep
	return nil
}
