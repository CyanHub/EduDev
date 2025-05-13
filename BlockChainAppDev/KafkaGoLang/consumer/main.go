package main

import (
	"BlockChainDev/config"
	"BlockChainDev/msg_datas"
	"BlockChainDev/pkg/kafka"
	"context"
	"fmt"
	"time"

	"github.com/IBM/sarama"
)

type KafkaConsumerGroupHandler struct {
}

func (kg *KafkaConsumerGroupHandler) Setup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (kg *KafkaConsumerGroupHandler) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

func NewKafkaTool() *kafka.KafkaTool {
	kt := &kafka.KafkaTool{}
	kafkaConfig := sarama.NewConfig()
	// 开启自动提交偏移量，设置提交间隔
	kafkaConfig.Consumer.Offsets.AutoCommit.Enable = true
	kafkaConfig.Consumer.Offsets.AutoCommit.Interval = 1 * time.Second
	return kt
}

func (kg *KafkaConsumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		fmt.Printf("消费者消费信息，消费的内容是: %s\n", string(message.Value))
		// 标记消息已消费
		session.MarkMessage(message, "")
		// 手动提交偏移量
		session.Commit()
	}
	return nil
}

func main() {
	config.Init()
	kt := kafka.NewKafkaTool()

	// 使用全新的消费者组 ID
	consumerGroup, err := sarama.NewConsumerGroup(kt.Brokers, "brand_new_consumer_group", kt.Config)
	if err != nil {
		panic(err)
	}
	defer consumerGroup.Close()

	ctx := context.Background()
	topics := []string{msg_datas.Topic_progress}
	handler := KafkaConsumerGroupHandler{}

	for {
		err := consumerGroup.Consume(ctx, topics, &handler)
		if err != nil {
			fmt.Printf("Error from consumer: %v\n", err)
		}
	}
}
