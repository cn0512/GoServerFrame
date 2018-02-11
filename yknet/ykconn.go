package yknet

import (
	"bufio"
	"fmt"
	//"errors"
	"io"
	"net"
	"time"

	cons "../ykconstant"
	yklog "../yklog"
	ykmsg "../ykmsg"
)

type ykConnect struct {
	Conn    net.Conn
	Uid     cons.UID
	Sendbuf chan []byte
	Close   chan interface{}
	Reader  *bufio.Reader
	Writer  *bufio.Writer
}

func (yk *ykConnect) Handle() {

	defer func() {
		yk.Conn.Close()
		close(yk.Close)
	}()

	go yk.Send()

	for {
		var header ykmsg.MsgHeader
		buf := header.Get()
		fmt.Println("header len=", len(buf))
		_, err := io.ReadFull(yk.Reader, buf)

		if err != nil {
			fmt.Println(err)
			yklog.Logout("client closed!")
			return
		}

		err = header.Set(buf)
		if err != nil {
			fmt.Println(err)
			yklog.Logout("client msg header is err!")
			return
		}
		body := make([]byte, header.MsgLen)
		_, err = io.ReadFull(yk.Reader, body)
		//设置在线映射
		yk.Uid = header.Uid
		ol.SetOL(&yk.Conn, header.Uid)
		//amqp 转发
		msg := ykmsg.RouterMsg{cons.Entity_Client, cons.Entity_Game, ykmsg.Msg{header, body}}
		mq_c2s.PushMsg(msg.Get())
	}
}

func (yk *ykConnect) Send() {
	for {

		select {
		case msg := <-yk.Sendbuf:
			//var header c.MsgHeader
			yk.Writer.Write(msg)
			yk.Writer.Flush()
		case <-yk.Close:
			ol.DelOL(yk.Uid)
			return
		case <-time.After(10 * time.Minute):
			yk.Conn.Close()
			yklog.Logout("connect timeout:%v", yk.Conn.RemoteAddr().String())
			return
		}
	}
}
