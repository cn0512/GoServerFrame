package main

import (
	"fmt"
	"os"
	"os/signal"

	mq "github.com/cn0512/GoServerFrame/Core/MQ/Redis"
)

var sub_topic = "MSvrs"

func main() {
	sub, err := mq.NewSub(sub_topic)
	defer sub.Shutdown()
	if err != nil {
		fmt.Println(err)
		return
	}
	//ctrl+C
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c
	fmt.Println("sub is finish")

}
