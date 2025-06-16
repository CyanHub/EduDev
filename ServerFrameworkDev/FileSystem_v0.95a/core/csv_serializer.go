package core

import (
	"FileSystem/global"
	"context"
	"errors"
	"reflect"
	"strings"

	"gorm.io/gorm/schema"
)
type CsvSerializer struct{
	schema.SerializerInterface
}



func (CsvSerializer) Scan(ctx context.Context, field *schema.Field, dst reflect.Value, dbValue interface{}) error{
	if dbValue == nil{
		return nil
	}
	var str string
	switch v := dbValue.(type){
		case string:
			str = v
		case []byte:
			str = string(v)
		default:
			return errors.New("不支持的数据类型")	
	}
	// str 
	values := strings.Split(str,global.CONFIG.MySQL.CsvSep)
	field.ReflectValueOf(ctx,dst).Set(reflect.ValueOf(values))

	return nil
}

func (CsvSerializer) Value(ctx context.Context, field *schema.Field, dst reflect.Value, fieldValue interface{}) (interface{}, error){
	if v,ok := fieldValue.([]string);ok{
		return strings.Join(v,global.CONFIG.MySQL.CsvSep),nil
	}
	return nil, errors.New("不支持的数据类型")
}

func RegisterCsvSerializer(){
	schema.RegisterSerializer("csv",CsvSerializer{})
}