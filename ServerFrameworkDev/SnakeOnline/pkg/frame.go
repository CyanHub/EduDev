package pkg

import (
	"encoding/binary"
	"io"
)

type MessageFrame struct {
	Length uint32 `json:"length,omitempty"`
	Type   uint8  `json:"type,omitempty"`
	Data   []byte `json:"data,omitempty"`
}

const (
	TypePlayerMove = iota + 1
	TypeFood
	TypeEat
)

func WriteFrame(writer io.Writer, msgType uint8, data []byte) error {
	length := uint32(len(data))
	err := binary.Write(writer, binary.BigEndian, length)
	if err != nil {
		return err
	}

	err = binary.Write(writer, binary.BigEndian, msgType)
	if err != nil {
		return err
	}

	_, err = writer.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func ReadFrame(reader io.Reader) (*MessageFrame, error) {
	messageFrame := &MessageFrame{}
	err := binary.Read(reader, binary.BigEndian, &messageFrame.Length)
	if err != nil {
		return nil, err
	}

	err = binary.Read(reader, binary.BigEndian, &messageFrame.Type)
	if err != nil {
		return nil, err
	}

	messageFrame.Data = make([]byte, messageFrame.Length)
	_, err = io.ReadFull(reader, messageFrame.Data)
	if err != nil {
		return nil, err
	}
	return messageFrame, nil
}
