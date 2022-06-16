package zdpgo_json

import "github.com/zhangdapeng520/zdpgo_json/query"

// 接口列表

// Write json写入文件的接口
type Write interface {
	// Dump 将Golang的任意obj对象，写入到filePath路径的json文件中
	Dump(filePath string, obj interface{}) error
}

// Read json文件读取的接口
type Read interface {
	// Load 将filePath路径的json文件，读取并反序列化到obj对象中
	Load(filePath string, obj interface{}) error
}

// ToJson 转换为json字符串的接口
type ToJson interface {
	// Dumps 将Golang的对象转换为json字符串
	Dumps(obj interface{}) (string, error)
}

// FromJson 处理json字符串的接口
type FromJson interface {
	// Loads 将json字符串转换为Go对象
	Loads(str string, obj interface{}) error
}

// Query 查询接口
type Query interface {
	// Get 从json字符串中查询指定数据
	Get(json, path string) query.Result
}
