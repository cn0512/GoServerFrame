package ykmq

import (
	"fmt"
	"log"

	c "../ykconstant"
	git "github.com/assembla/cony"
)

type Consumer struct {
	Client   *git.Client
	Consumer *git.Consumer
	Mq       chan []byte
}

func (yk *Consumer) PushMsg(buf []byte) {
}

func (yk *Consumer) PopMsg() {

	for yk.Client.Loop() {
		select {
		case msg := <-yk.Consumer.Deliveries():
			log.Printf("Consumer pop msg: %q\n", msg.Body)
			// If when we built the consumer we didn't use
			// the "git.AutoAck()" option this is where we'd
			// have to call the "amqp.Deliveries" methods "Ack",
			// "Nack", "Reject"
			//
			// msg.Ack(false)
			// msg.Nack(false)
			// msg.Reject(false)
			//make response msg 2 mq 2 client
			yk.Mq <- msg.Body
		case err := <-yk.Consumer.Errors():
			fmt.Printf("Consumer error: %v\n", err)
		case err := <-yk.Client.Errors():
			fmt.Printf("Client error: %v\n", err)
		}
	}
}

func (yk *Consumer) Create(url string, s2c bool) {
	cli := git.NewClient(
		git.URL(url),
		git.Backoff(git.DefaultBackoff),
	)
	var que *git.Queue
	if s2c {
		que = &git.Queue{
			AutoDelete: true,
			Name:       c.AMQP_queueName2,
		}
	} else {

		que = &git.Queue{
			AutoDelete: true,
			Name:       c.AMQP_queueName,
		}
	}
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
	var bnd git.Binding
	if s2c {

		bnd = git.Binding{
			Queue:    que,
			Exchange: exc,
			Key:      c.AMQP_routingKey2,
		}
	} else {

		bnd = git.Binding{
			Queue:    que,
			Exchange: exc,
			Key:      c.AMQP_routingKey,
		}
	}
	cli.Declare([]git.Declaration{
		git.DeclareQueue(que),
		git.DeclareExchange(exc),
		git.DeclareBinding(bnd),
	})

	// Declare and register a consumer
	cns := git.NewConsumer(
		que,
		git.AutoAck(), // Auto sign the deliveries
	)
	cli.Consume(cns)

	yk.Client = cli
	yk.Consumer = cns

	go yk.PopMsg()
}
