package initialize

import (
	"ServerFramework/global"
	"context"
	"encoding/json"
	"log"

	"gorm.io/gorm"
)

// RegisterSerializer 注册序列化器
// 序列化器用于在数据库操作前将结构体中的 JSON 字段序列化为字符串，
// 在数据库操作后将字符串反序列化为 JSON 字段。
func 
RegisterSerializer() {
	if global.DB == nil {
		log.Fatal("数据库连接未初始化，无法注册序列化器")
	}
	global.DB.Callback().Create().Before("gorm:before_create").Register("json_serializer:before_create", jsonSerializer)
	global.DB.Callback().Update().Before("gorm:before_update").Register("json_serializer:before_update", jsonSerializer)
	global.DB.Callback().Query().After("gorm:after_query").Register("json_serializer:after_query", jsonDeserializer)
}

// jsonSerializer 序列化 JSON 字段
func jsonSerializer(db *gorm.DB) {
	if db.Statement.Schema != nil {
		for _, field := range db.Statement.Schema.Fields {
			if field.Tag.Get("serializer") == "json" {
				value, ok := field.ValueOf(context.Background(), db.Statement.ReflectValue)
				if ok {
					bytes, err := json.Marshal(value)
					if err == nil {
						err := field.Set(context.Background(), db.Statement.ReflectValue, string(bytes))
						if err == nil {
							// 设置成功
						}
					}
				}
			}
		}
	}
}

// jsonDeserializer 反序列化 JSON 字段
func jsonDeserializer(db *gorm.DB) {
	if db.Statement.Schema != nil {
		for _, field := range db.Statement.Schema.Fields {
			if field.Tag.Get("serializer") == "json" {
				value, ok := field.ValueOf(context.Background(), db.Statement.ReflectValue)
				if ok {
					str, ok := value.(string)
					if ok {
						var temp interface{}
						err := json.Unmarshal([]byte(str), &temp)
						if err == nil {
							err := field.Set(context.Background(), db.Statement.ReflectValue, temp)
							if err == nil {
								// 设置成功
							}
						}
					}
				}
			}
		}
	}
}
