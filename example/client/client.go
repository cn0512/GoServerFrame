package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/signal"
	"time"

	cons "../../ykconstant"
	"../../yklog"
	//mq "../../ykmq"
	ykmsg "../../ykmsg"
	//yknet "../../yknet"
)

const (
	C_SLEEP = 1000 * 1000 * 1000 * 5
)

func send(conn net.Conn) {

	defer conn.Close()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "exit" {
			os.Exit(0)
		}
		fmt.Println("input=", line) // Println will add back the final '\n'

		bufbody := []byte(line)
		var msgHeader ykmsg.MsgHeader
		var ykgob ykmsg.GobMsg
		bufMsg, err := ykgob.Encode(&bufbody, nil)
		msgHeader.MsgID = 1000
		msgHeader.Uid = 9527
		msgHeader.MsgLen = (int32)(len(bufMsg))
		fmt.Println(msgHeader)
		bufHeader := msgHeader.Get()
		buf := append(bufHeader, bufMsg...)
		if err != nil {
			fmt.Println(err)
		}

		conn.Write(buf)
		/*decode*/
		fmt.Println(buf)
		var msg []byte = make([]byte, len(bufMsg))
		err = ykgob.Decode(bufMsg, &msg)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(msg)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func recv(conn net.Conn) {
	for {
		var buf []byte
		len, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		if len > 0 {
			fmt.Println("recv=", buf)
			var ykgob ykmsg.GobMsg
			var msg ykmsg.Msg
			ykgob.Decode(buf, msg)
			fmt.Println(msg)
		}

	}
}

func create_client(addr string) {
	//addr := "127.0.0.1:7056"
	for {
		conn, err := net.Dial("tcp", addr)
		if err == nil {
			go send(conn)
			go recv(conn)
			return
		}
		yklog.Logout("connect to %v error: %v", addr, err)

		time.Sleep(C_SLEEP)
		continue
	}

}

func main() {

	//p := mq.Consumer{nil, nil, make(chan []byte, 3)}
	//p.Create(c.AMQP_uri)
	create_client(cons.TCP_IP)

	// close
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	sig := <-c
	yklog.Logout("client quit!(signal: %v)", sig)
}
