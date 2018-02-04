package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/signal"
	"time"

	"../../yklog"
)

const (
	C_SLEEP = 1000 * 1000 * 1000 * 5
)

func handle_client(conn net.Conn) {

	defer conn.Close()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "exit" {
			os.Exit(0)
		}
		fmt.Println(line) // Println will add back the final '\n'
		conn.Write([]byte(line))
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func create_client(addr string) {
	//addr := "127.0.0.1:7056"
	for {
		conn, err := net.Dial("tcp", addr)
		if err == nil {
			go handle_client(conn)
			return
		}
		yklog.Logout("connect to %v error: %v", addr, err)

		time.Sleep(C_SLEEP)
		continue
	}

}

func main() {
	create_client("127.0.0.1:6018")

	// close
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	sig := <-c
	yklog.Logout("client quit!(signal: %v)", sig)
}
