package ykmq

import (
	"fmt"

	"github.com/streadway/amqp"
)

type YKMQ struct {
	conn            *amqp.Connection
	ch              *amqp.Channel
	passiveExchange bool
}

func Create(uri string, pex bool) (*YKMQ, error) {
	conn, ch, err := createConnAndChan(uri)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &YKMQ{conn: conn, ch: ch, passiveExchange: pex}, nil
}

func createConnAndChan(dsn string) (*amqp.Connection, *amqp.Channel, error) {
	conn, err := amqp.Dial(dsn)
	if err != nil {
		return nil, nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, nil, err
	}

	return conn, ch, nil
}

func (yk *YKMQ) Publish(exchange, key string, opts amqp.Publishing) error {
	return yk.ch.Publish(exchange, key, false, false, opts)
}

func (yk *YKMQ) CreateConsumer(exchange, key, kind, queue string, durable bool) (<-chan amqp.Delivery, error) {
	if err := yk.WithExchange(exchange, kind, durable); err != nil {
		return nil, err
	}

	q, err := yk.ch.QueueDeclare(queue, durable, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	if err := yk.ch.QueueBind(q.Name, key, exchange, false, nil); err != nil {
		return nil, err
	}

	return yk.ch.Consume(q.Name, "", false, false, false, false, nil)
}

func (yk *YKMQ) WithExchange(exchange, kind string, durable bool) error {
	if yk.passiveExchange {
		return yk.ch.ExchangeDeclarePassive(exchange, kind, durable, false, false, false, nil)
	}

	return yk.ch.ExchangeDeclare(exchange, kind, durable, false, false, false, nil)
}

func (yk *YKMQ) WithQos(count, size int, global bool) error {
	return yk.ch.Qos(count, size, global)
}

func (yk *YKMQ) NotifyClose(c chan *amqp.Error) chan *amqp.Error {
	return yk.conn.NotifyClose(c)
}

func (yk *YKMQ) Close() error {
	if err := yk.ch.Close(); err != nil {
		return err
	}

	if yk.conn != nil {
		return yk.conn.Close()
	}

	return nil
}
