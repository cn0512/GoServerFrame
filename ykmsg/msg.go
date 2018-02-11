package ykmsg

import (
	"errors"

	cons "../ykconstant"
)

type MsgHeader struct {
	MsgLen int32
	MsgID  int32
	Uid    cons.UID
}

type Msg struct {
	MsgHeader
	Body []byte
}

func (yk *MsgHeader) GetLen() int32 {
	return 12
}

func (yk *MsgHeader) Get() (buf []byte) {

	buf1 := cons.IntToBytes(yk.MsgLen)
	buf2 := cons.IntToBytes(yk.MsgID)
	buf3 := cons.IntToBytes(int32(yk.Uid))

	buf4 := append(buf1, buf2...)
	buf = append(buf4, buf3...)

	return buf
}

func (yk *MsgHeader) Set(buf []byte) error {

	s := len(buf) / 3
	if s < 4 {
		return errors.New("header len is little")
	}
	yk.MsgLen = cons.BytesToInt(buf[:s])
	yk.MsgID = cons.BytesToInt(buf[s : 2*s])
	yk.Uid = cons.UID(cons.BytesToInt(buf[2*s:]))
	//fmt.Println(yk)
	return nil

}

type RouterMsg struct {
	SrcEntity int32 //源Entity
	DstEntity int32 //目标Entity
	InnerMsg  Msg
}

func (yk *RouterMsg) Get() (buf []byte) {
	/*
		buf1 := cons.IntToBytes(yk.SrcEntity)
		buf2 := cons.IntToBytes(yk.DstEntity)
		buf3 := yk.InnerMsg.Get()

		buf4 := append(buf1, buf2...)
		buf = append(buf4, buf3...)
	*/
	var ykgob GobMsg
	buf, _ = ykgob.Encode(yk, nil)
	return buf
}

func (yk *RouterMsg) Set(buf []byte) error {
	var ykgob GobMsg
	return ykgob.Decode(buf, yk)
}
