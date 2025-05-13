package kafka

import (
	"BlockChainDev/config"
	"context"
	"fmt"
	"time"

	"github.com/IBM/sarama"
)

type KafkaProducer struct {
	Producer sarama.SyncProducer
}

type KafkaConsumer struct {
	Consumer      sarama.Consumer
	ConsumerGroup sarama.ConsumerGroup
}

type KafkaTool struct {
	Brokers       []string `json:"brokers,omitempty"`
	Config        *sarama.Config
	KafkaProducer *KafkaProducer
	KafkaConsumer *KafkaConsumer
}

func NewKafkaTool() *KafkaTool {
	kt := &KafkaTool{}
	brokers := make([]string, 0)
	brokers = append(brokers, config.CONFIG.Kafka.Host+":"+config.CONFIG.Kafka.Port)
	kt.Brokers = brokers

	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Producer.RequiredAcks = sarama.WaitForAll
	kafkaConfig.Producer.Compression = sarama.CompressionSnappy
	kafkaConfig.Producer.Partitioner = sarama.NewRandomPartitioner
	kafkaConfig.Producer.Return.Successes = true
	kafkaConfig.Producer.Return.Errors = true
	// 从最新的消息开始消费
	kafkaConfig.Consumer.Offsets.Initial = sarama.OffsetNewest
	// 开启自动提交偏移量，设置提交间隔
	kafkaConfig.Consumer.Offsets.AutoCommit.Enable = true
	kafkaConfig.Consumer.Offsets.AutoCommit.Interval = 1 * time.Second
	kt.Config = kafkaConfig

	// Create a KafkaProducer instance
	kafkaProducer := &KafkaProducer{}
	producer, err := sarama.NewSyncProducer(brokers, kafkaConfig)
	if err != nil {
		panic(err)
	}
	kafkaProducer.Producer = producer
	// 将KAFKAPRODUCER实例分配到Kafkatool结构
	kt.KafkaProducer = kafkaProducer

	// 创建一个kafkaconsumer实例
	kafkaConsumer := &KafkaConsumer{}
	consumer, err := sarama.NewConsumer(brokers, kafkaConfig)
	if err != nil {
		panic(err)
	}
	kafkaConsumer.Consumer = consumer
	// 将kafkaconsumer实例分配到kafkatool struct
	kt.KafkaConsumer = kafkaConsumer

	return kt
}

func (kt *KafkaTool) AddConsumeHandler(groupID string, topics []string, handler sarama.ConsumerGroupHandler) {
	consumerGroup, err := sarama.NewConsumerGroup(kt.Brokers, groupID, kt.Config)
	if err != nil {
		panic(err)
	}
	kt.KafkaConsumer.ConsumerGroup = consumerGroup
	for {
		err = consumerGroup.Consume(context.Background(), topics, handler)
		if err != nil {
			fmt.Println("消息信息错误：", err)
		}
	}
}
func (kt *KafkaTool) ProduceMsg(topic string, message string) {
	saramaMessage := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message), // 将消息内容编码为字符串
	}
	// 使用生产者发送消息并获取消息的分区和偏移
	partition, offset, err := kt.KafkaProducer.Producer.SendMessage(saramaMessage)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Produced message:", saramaMessage)
	// 输出成功发送的分区和偏移信息
	fmt.Printf("Partition: %d, Offset: %d\n", partition, offset)
}
