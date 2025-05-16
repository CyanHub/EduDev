package common

import (
    "encoding/binary"
    "io"
)

// MessageFrame 消息帧结构
type MessageFrame struct {
    Length uint32 // 消息长度
    Type   uint8  // 消息类型
    Data   []byte // 消息内容
}

const (
    TypePlayerMove = iota + 1
    TypeFood
    TypeEat
	TypePlayerLeft
)

// 写入消息帧
func WriteFrame(writer io.Writer, msgType uint8, data []byte) error {
    length := uint32(len(data))
    
    // 写入长度
    if err := binary.Write(writer, binary.BigEndian, length); err != nil {
        return err
    }
    
    // 写入类型
    if err := binary.Write(writer, binary.BigEndian, msgType); err != nil {
        return err
    }
    
    // 写入数据
    _, err := writer.Write(data)
    return err
}

// 读取消息帧
func ReadFrame(reader io.Reader) (*MessageFrame, error) {
    frame := &MessageFrame{}
    
    // 读取长度
    if err := binary.Read(reader, binary.BigEndian, &frame.Length); err != nil {
        return nil, err
    }
    
    // 读取类型
    if err := binary.Read(reader, binary.BigEndian, &frame.Type); err != nil {
        return nil, err
    }
    
    // 读取数据
    frame.Data = make([]byte, frame.Length)
    _, err := io.ReadFull(reader, frame.Data)
    if err != nil {
        return nil, err
    }
    
    
    return frame, nil
}