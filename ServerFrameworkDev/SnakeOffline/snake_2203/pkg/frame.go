package pkg

import (
	"encoding/binary"
	"io"
)

type FrameMessage struct {
	Length uint32 `json:"length"`
	Type   uint8  `json:"type"`
	Data   []byte `json:"data"`
}

const (
	TypePlayerMove = iota + 1
	TypeFood
	TypeEat
)

func WriteFrame(writer io.Writer, msgType uint8, data []byte) error {
	length := uint32(len(data))

	// 写入固定大小的消息长度
	err := binary.Write(writer, binary.BigEndian, length)
	if err != nil {
		return err
	}

	// 写入固定大小的消息类型
	err = binary.Write(writer, binary.BigEndian, msgType)
	if err != nil {
		return err
	}

	// 写入变长的消息内容
	_, err = writer.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func ReadFrame(reader io.Reader) (*FrameMessage, error) {
	messageFrame := FrameMessage{}

	// 读取固定大小的消息长度
	err := binary.Read(reader, binary.BigEndian, &messageFrame.Length)
	if err != nil {
		return nil, err
	}

	// 读取固定大小的消息类型
	err = binary.Read(reader, binary.BigEndian, &messageFrame.Type)
	if err != nil {
		return nil, err
	}

	// 读取消息的具体内容
	messageFrame.Data = make([]byte, messageFrame.Length)
	//buffer := make([]byte, 2048)
	_, err = io.ReadFull(reader, messageFrame.Data)
	if err != nil {
		return nil, err
	}

	return &messageFrame, nil
}
