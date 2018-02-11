package GamePlus

/*
	a simple game :compare one single cards who is the biger
*/

import (
	"fmt"

	"github.com/golang/protobuf/proto"

	cons "../../ykconstant"
	mq "../../ykmq"
	ykmsg "../../ykmsg"
	"../../ykprotoco"
)

const (
	MAX_GAME_PLAYER = 8
	MAX_CARDS_NUM   = 1

	Info = "GamePlus_Logic1 a simple game with 2 players"

	AMQP_Queue = "GamePlus_Logic1"
)

type GamePlus_Logic1 struct {
	players []*YKGameMsg.StPlayer

	Consumer *mq.Consumer
	Producer *mq.Producer
}

func (yk *GamePlus_Logic1) Init() {
	yk.players = make([]*YKGameMsg.StPlayer, MAX_GAME_PLAYER)

	yk.Consumer = &mq.Consumer{nil, nil, make(chan []byte, cons.Buf_Size)}
	yk.Consumer.Create(cons.AMQP_uri, false)

	yk.Producer = &mq.Producer{nil, nil, make(chan []byte, cons.Buf_Size)}
	yk.Producer.Create(cons.AMQP_uri, true)
}

func (yk *GamePlus_Logic1) Run() {
	//从AMQP拉取消息
	go yk.Consumer.PopMsg()
	for {

		select {
		case msg := <-yk.Consumer.Mq:

			//解码Gate消息
			var r ykmsg.RouterMsg
			var gobmsg ykmsg.GobMsg
			err := gobmsg.Decode(msg, &r)
			if err != nil {
				fmt.Println(err)
				return
			} else {
				fmt.Println("gameplus1 mq s2c=", r)
			}
			if r.DstEntity != cons.Entity_Game {
				fmt.Println("gameplus1 mq entity=", r.DstEntity)
			}
			//处理Msg
			//echo response 2 client
			var resp ykmsg.RouterMsg
			resp.SrcEntity = cons.Entity_Game
			resp.DstEntity = cons.Entity_Client
			buf, _ := gobmsg.Encode(&resp, nil)
			resp.InnerMsg = r.InnerMsg
			yk.Producer.PushMsg(buf)
		}
	}
}

func (yk *GamePlus_Logic1) End() {

}

func (yk *GamePlus_Logic1) Ready([]byte) {

}

func (yk *GamePlus_Logic1) About() string {
	return Info
}

func (yk *GamePlus_Logic1) GameBegin(buf []byte) {
	fmt.Println(buf)
	var msg YKGameMsg.GameBegin
	err := proto.Unmarshal(buf, &msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	yk.players = msg.Players
	fmt.Println(yk)
}
