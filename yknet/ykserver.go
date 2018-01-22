package yknet

import (
	"net"

	"../yklog"
)

func handle_server(conn net.Conn) {

	buf := make([]byte, 50)

	defer conn.Close()

	for {
		n, err := conn.Read(buf)

		if err != nil {
			yklog.Logout("client closed!")
			return
		}
		yklog.Logout("recv msg:%v", string(buf[0:n]))
	}
}

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
		go handle_server(conn)
	}
}
