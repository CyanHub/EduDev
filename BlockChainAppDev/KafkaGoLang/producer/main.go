package main

import (
	"BlockChainDev/config"
	"BlockChainDev/msg_datas"
	"BlockChainDev/pkg/kafka"
	"fmt"
)

func main() {
	config.Init()
	kt := kafka.NewKafkaTool()
	for i := 0; i < 10; i++ {
		kt.ProduceMsg(msg_datas.Topic_progress, fmt.Sprintf("progress:%d", i+1))

	}
}
