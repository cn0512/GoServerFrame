package main

import (
	cfg "MSvrs/Config"
	//mysqlx "MSvrs/Core/DB/Mysqlx"
	RPC "MSvrs/Core/RPC"
	"context"
	"fmt"
	//"github.com/jmoiron/sqlx"
	ykmq "MSvrs/Core/MQ/YKMQ"
	tl "MSvrs/Core/Utils"
	"time"
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
