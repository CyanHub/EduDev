package msg_datas

// const Topic_hello string = "topic_hello"
const Topic_progress string = "progress"

type BaseMsg struct {
	MsgId   int    `json:"msg_id,omitempty"`
	MsgData string `json:"msg_data"`
}

type ProduceMsg struct {
	ToUid   string `json:"to_uid,omitempty"`
	SendUid string `json:"send_uid,omitempty"`
	MsgType int    `json:"msg_type,omitempty" ` // 消息类型
	Content string `json:"content,omitempty"`
}
