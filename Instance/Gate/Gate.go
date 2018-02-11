package main

/*
	网关：
	（1）接受client连接
	（2）转发协议到amqp
	（3）验证协议密钥
*/

import (
	"flag"
	"fmt"

	"os"
	"os/signal"

	c "../../ykconstant"
	"../../yknet"
)

var (
	version = flag.String("v", "0.0.1", "version = ?")
	tcp_ip  = flag.String("ip", "127.0.0.1:6018", "ip:port = ?")
)

func init() {
	flag.Parse()
}

func main() {
	fmt.Println("version,ip=", *version, *tcp_ip)
	fmt.Println("amqp=", c.AMQP_uri)

	yknet.Create_server(c.TCP_IP)

	q := make(chan os.Signal, 1)
	signal.Notify(q, os.Interrupt, os.Kill)
	<-q
}
