package ykmq

import (
	"fmt"

	git "github.com/assembla/cony"
	"github.com/streadway/amqp"

	c "../ykconstant"
)

type Producer struct {
	Client    *git.Client
	Publisher *git.Publisher
	Mq        chan []byte
}

func (yk *Producer) PushMsg(buf []byte) {
	fmt.Printf("Producer publishing\n")

	err := yk.Publisher.Publish(amqp.Publishing{
		Body: buf,
	})
	if err != nil {
		fmt.Printf("Producer publish error: %v\n", err)
	}
}

func (yk *Producer) PopMsg() {

	for yk.Client.Loop() {
		select {
		case err := <-yk.Client.Errors():
			fmt.Printf("Client error: %v\n", err)
		case blocked := <-yk.Client.Blocking():
			fmt.Printf("Client is blocked %v\n", blocked)
		}
	}
}

func (yk *Producer) Create(url string, s2c bool) {

	cli := git.NewClient(
		git.URL(url),
		git.Backoff(git.DefaultBackoff),
	)

	// Declare the exchange we'll be using
	var exc git.Exchange
	if s2c {

		exc = git.Exchange{
			Name:       c.AMQP_exchangeName2,
			Kind:       c.AMQP_exchangeType2,
			AutoDelete: true,
		}
	} else {

		exc = git.Exchange{
			Name:       c.AMQP_exchangeName,
			Kind:       c.AMQP_exchangeType,
			AutoDelete: true,
		}
	}
	cli.Declare([]git.Declaration{
		git.DeclareExchange(exc),
	})

	// Declare and register a publisher
	// with the git client
	var pbl *git.Publisher
	if s2c {
		pbl = git.NewPublisher(exc.Name, c.AMQP_routingKey2)
	} else {
		pbl = git.NewPublisher(exc.Name, c.AMQP_routingKey)
	}

	cli.Publish(pbl)
	// Launch a go routine and publish a message.
	// "Publish" is a blocking method this is why it
	// needs to be called in its own go routine.
	//

	yk.Client = cli
	yk.Publisher = pbl

	// Client loop sends out declarations(exchanges, queues, bindings
	// etc) to the AMQP server. It also handles reconnecting.
	go yk.PopMsg()
}
