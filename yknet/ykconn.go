package yknet

import (
	"bufio"
	"errors"
	"io"
	"net"
	"time"

	c "../ykconstant"
	yklog "../yklog"
)

type MsgHeader struct {
	MsgLen int32
	MsgID  int32
}

type Msg struct {
	MsgHeader
	body []byte
}

func (yk *MsgHeader) GetLen() int32 {
	return 8
}

func (yk *MsgHeader) Get() (buf []byte, msglen int32) {

	buf1 := c.IntToBytes(yk.MsgLen)
	buf2 := c.IntToBytes(yk.MsgID)

	buf = append(buf1, buf2...)
	msglen = int32(len(buf))
	return buf, msglen
}

func (yk *MsgHeader) Set(buf []byte) error {
	s := len(buf) / 2
	if s < 4 {
		return errors.New("header len is little")
	}
	yk.MsgLen = c.BytesToInt(buf[:s])
	yk.MsgID = c.BytesToInt(buf[s:])
	return nil
}

type ykConnect struct {
	Conn net.Conn

	Sendbuf chan []byte
	Close   chan interface{}
	Reader  *bufio.Reader
	Writer  *bufio.Writer
}

func (yk *ykConnect) Handle() {

	defer func() {
		yk.Conn.Close()
		close(yk.Close)
	}()

	go yk.Send()

	for {
		var header MsgHeader
		buf, _ := header.Get()

		_, err := io.ReadFull(yk.Reader, buf)

		if err != nil {
			yklog.Logout("client closed!")
			return
		}

		err = header.Set(buf)
		if err == nil {
			yklog.Logout("client msg header is err!")
			return
		}
		body := make([]byte, header.MsgLen)
		_, err = io.ReadFull(yk.Reader, body)
		//amqp 转发
	}
}

func (yk *ykConnect) Send() {
	for {

		select {
		case msg := <-yk.Sendbuf:
			//var header c.MsgHeader
			yk.Writer.Write(msg)
			yk.Writer.Flush()
		case <-yk.Close:
			return
		case <-time.After(20 * time.Second):
			yk.Conn.Close()
			yklog.Logout("connect timeout:%v", yk.Conn.RemoteAddr().String())
			return
		}
	}
}
