// yk
package main

import (
	"os"
	"os/signal"

	"../../yklog"

	"../../yknet"
)

func main() {

	yklog.Logout("yk frame start!")

	yknet.Create_server("127.0.0.1:7056")

	// close
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	sig := <-c
	yklog.Logout("server quit!(signal: %v)", sig)
}
