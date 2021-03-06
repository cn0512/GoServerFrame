package main

import (
	"bufio"
	"os"
	"os/signal"
	_ "strconv"

	mq "github.com/cn0512/GoServerFrame/Core/MQ/Redis"
	"golang.org/x/fmt"
)

var pub_topic = "MSvrs"

func main() {
	pub := mq.NewPub()
	defer pub.Shutdown()
	//pub.Publish(pub_topic,[]byte(strconv.Itoa(1000)))

	//use input to pub
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "exit" {
			os.Exit(0)
		}
		pub.Publish(pub_topic, []byte(string(line)))
	}
	//ctrl+C
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c
	fmt.Println("pub is finish")
}
