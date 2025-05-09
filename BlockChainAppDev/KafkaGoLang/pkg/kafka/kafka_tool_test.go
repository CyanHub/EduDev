package kafka_tool_test

type ProduceMsg_test struct {
    ToUid   string `json:"to_uid" bson:"to_uid"`     // 接收者id
    SendUid string `json:"send_uid" bson:"send_uid"` // 发送者id
    MsgType int    `json:"msg_type" bson:"msg_type"` // 消息类型

}