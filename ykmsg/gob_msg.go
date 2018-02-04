package ykmsg

import (
	"bytes"
	"encoding/gob"
)

type GobMsg struct{}

func (mp GobMsg) Encode(msg interface{}, buf []byte) ([]byte, error) {

	buffer := bytes.NewBuffer(buf)

	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(msg)
	if err != nil {
		return buf, err
	}
	buf = buffer.Bytes()
	return buf, nil
}

func (mp GobMsg) Decode(data []byte, msg interface{}) error {
	decoder := gob.NewDecoder(bytes.NewBuffer(data))
	err := decoder.Decode(msg)
	return err
}
