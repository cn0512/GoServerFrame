package yknet

import (
	"bufio"
	"fmt"
	"net"

	cons "../ykconstant"
	"../yklog"

	mq "../ykmq"

	storge "../ykredis"
)

var ol storge.UserOL
var mq_c2s *mq.Producer
var mq_s2c *mq.Consumer

func Create_server(addr string) {
	//addr := "127.0.0.1:7056"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		yklog.Logout(err.Error())
		return
	}
	fmt.Println("listen =", addr)
	defer listener.Close()

	mq_c2s = &mq.Producer{nil, nil, make(chan []byte, cons.Buf_Size)}
	mq_c2s.Create(cons.AMQP_uri, false)

	mq_s2c = &mq.Consumer{nil, nil, make(chan []byte, cons.Buf_Size)}
	mq_s2c.Create(cons.AMQP_uri, true)

	ol.Create()

	go func() {
		for {
			select {
			case msg := <-mq_s2c.Mq:
				fmt.Println("yknet s2c msg=", msg)

			}
		}
	}()

	for {
		conn, err := listener.Accept()
		if err != nil {
			yklog.Logout(err.Error())
		}
		Conn := &ykConnect{conn, 0,
			make(chan []byte, 2),
			make(chan interface{}),
			bufio.NewReaderSize(conn, cons.Buf_Size),
			bufio.NewWriterSize(conn, cons.Buf_Size)}
		go Conn.Handle()
	}
}
