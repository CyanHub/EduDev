package msg_datas

// const Topic_hello string = "topic_hello"
// Topic_progress 定义了 Kafka 主题名称，用于表示进度相关消息的主题
const Topic_progress string = "progress"

// BaseMsg 表示基础消息结构体，包含消息的基本信息
type BaseMsg struct {
	// MsgId 消息的唯一标识符，omitempty 表示当该字段为空时，在 JSON 序列化时忽略该字段
	MsgId int `json:"msg_id,omitempty"`
	// MsgData 消息的具体内容，在 JSON 序列化时固定为 "msg_data" 字段名
	MsgData string `json:"msg_data"`
}

// ProduceMsg 表示用于生产发送的消息结构体，包含消息的接收者、发送者、类型和内容等信息
type ProduceMsg struct {
	// ToUid 消息的接收用户 ID，omitempty 表示当该字段为空时，在 JSON 序列化时忽略该字段
	ToUid string `json:"to_uid,omitempty"`
	// SendUid 消息的发送用户 ID，omitempty 表示当该字段为空时，在 JSON 序列化时忽略该字段
	SendUid string `json:"send_uid,omitempty"`
	// MsgType 消息的类型，omitempty 表示当该字段为空时，在 JSON 序列化时忽略该字段
	MsgType int `json:"msg_type,omitempty" ` // 消息类型
	// Content 消息的具体内容，omitempty 表示当该字段为空时，在 JSON 序列化时忽略该字段
	Content string `json:"content,omitempty"`
}
