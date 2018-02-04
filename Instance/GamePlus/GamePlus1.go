package GamePlus

/*
	a simple game :compare one single cards who is the biger
*/

import (
	"fmt"

	"github.com/golang/protobuf/proto"

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
}

func (yk *GamePlus_Logic1) Init() {
	yk.players = make([]*YKGameMsg.StPlayer, MAX_GAME_PLAYER)
}

func (yk *GamePlus_Logic1) Run() {
	//从AMQP拉取消息
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
