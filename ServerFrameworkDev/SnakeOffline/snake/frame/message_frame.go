package frame

import (
	"encoding/binary"
	"io"
)

const (
	TypePlayerMove = iota + 1
	TypeFood
	TypeEat
	TypePlayerLeft
)

type MessageFrame struct {
	Length uint32 `json:"length"`
	Type   uint8  `json:"type"`
	Data   []byte `json:"data"`
}

func WriteFrame(writer io.Writer, msgType uint8, data []byte) (int, error) {
	length := uint32(len(data))
	err := binary.Write(writer, binary.BigEndian, length)
	if err != nil {
		return 0, err
	}
	err = binary.Write(writer, binary.BigEndian, msgType)
	if err != nil {
		return 0, err
	}
	return writer.Write(data)
}

func ReadFrame(reader io.Reader) (*MessageFrame, error) {
	message := MessageFrame{}

	err := binary.Read(reader, binary.BigEndian, &message.Length)
	if err != nil {
		return nil, err
	}
	err = binary.Read(reader, binary.BigEndian, &message.Type)
	if err != nil {
		return nil, err
	}

	buffer := make([]byte, message.Length)
	_, err = io.ReadFull(reader, buffer)
	if err != nil {
		return nil, err
	}
	message.Data = buffer
	return &message, nil
}
