package kafka_tool_test

import (
	// "fmt"

	"github.com/IBM/sarama"
)

type KafkaProducer struct {
	Producer sarama.SyncProducer // kafka生产者
}

type KafkaConsumer struct {
	Consumer             sarama.Consumer             // kafka消费者
	ConsumerGroup        sarama.ConsumerGroup        // kafka消费者组
	ConsumerGroupSession sarama.ConsumerGroupSession // kafka消费者组会话
	ConsumerGroupClaim   sarama.ConsumerGroupClaim   // kafka消费者组声明
}

type KafkaTool struct {
	Brokers       []string       `json:"brokers,omitempty"` // broker地址列表，多个地址用逗号分隔
	Config        sarama.Config  // kafka配置项
	KafkaProducer *KafkaProducer // kafka生产者
	KafkaConsumer *KafkaConsumer // kafka消费者
}

// func NewKafkaTool() {
// 	kt := &KafkaTool{}
// 	brokers := make([]string, 0) // broker地址列表，多个地址用逗号分隔
// 	brokers = append(brokers, config.CONFIG.Kafka.Host+":"+config.CONFIG.Kafka.Port)
// 	kt.Brokers = brokers

// 	kafkaConfig := sarama.NewConfig()                              // kafka配置项
// 	kafkaConfig.Producer.RequiredAcks = sarama.WaitForAll          // 等待所有副本都收到消息
// 	kafkaConfig.Producer.Compression = sarama.CompressionSnappy    // 压缩消息
// 	kafkaConfig.Producer.Partitioner = sarama.NewRandomPartitioner // 随机分区
// 	kafkaConfig.Producer.Return.Successes = true                   // 成功返回
// 	kafkaConfig.Producer.Return.Errors = true                      // 失败返回
// 	kafkaConfig.Consumer.Return.Errors = true                      // 失败返回
// 	kt.Config = *kafkaConfig                                       // kafka配置项

// 	KafkaProducer := &KafkaProducer{} // kafka生产者
// 	producer, err := sarama.NewSyncProducer(Brokers, kafkaConfig)
// 	if err != nil {
// 		panic(err)
// 	}
// 	KafkaConsumer.Consumer = consumer
// 	return kt
// }

// func (kt *KafkaTool) AddConsumeHandler(groupID string, topics []string, handler sarama.ConsumerGroupHandler) error {
// 	// 创建消费者组
// 	consumerGroup, err := sarama.NewConsumerGroup(kt.Brokers, groupID, &kt.Config)
// 	if err != nil {
// 		panic(err)
// 	}
// 	kt.KafkaConsumer.ConsumerGroup = consumerGroup

// 	// 启动消费者组
// 	for {
// 		err := consumerGroup.Consume(context.Background(), topics, handler)
// 		if err != nil {
// 			fmt.Println("Error from consumer:", err)
// 		}
// 	}
// }

// func (kt *KafkaTool) ProduceMsg(topic string, msg string) error {
// 	saramaMessage := &sarama.ProducerMessage{ // 消息结构
// 		Topic: topic, // 消息主题
// 		Value: sarama.StringEncoder(msg), // 消息具体内容
// 	}
// 	partition, offset, err := kt.KafkaProducer.Producer.SendMessage(saramaMessage) // 发送消息
// 	if err!= nil {
// 		fmt.Println("发送消息失败:", err)
// 		return 
// 	}
// 	fmt.Printf("消息发送成功，分区: %d, 偏移量: %d\n", partition, offset)
// }
