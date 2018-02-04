package yknet

import (
	"bufio"
	"net"

	c "../ykconstant"
	"../yklog"
)

func Create_server(addr string) {
	//addr := "127.0.0.1:7056"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		yklog.Logout(err.Error())
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			yklog.Logout(err.Error())
		}
		Conn := &ykConnect{conn,
			make(chan []byte, 2),
			make(chan interface{}),
			bufio.NewReaderSize(conn, c.Buf_Size),
			bufio.NewWriterSize(conn, c.Buf_Size)}
		go Conn.Handle()
	}
}
