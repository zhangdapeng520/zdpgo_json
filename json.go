package zdpgo_json

import (
	"encoding/json"
	"os"

	"github.com/zhangdapeng520/zdpgo_json/query"
)

// Json 处理json的核心对象
type Json struct {
	Query *query.Query // 查询核心对象

	// 方法列表
	Dump  func(filePath string, obj interface{}) error
	Load  func(filePath string, obj interface{}) error
	Dumps func(obj interface{}) (string, error)
	Loads func(str string, obj interface{}) error
}

// New 创建新的处理json的对象示例
func New() *Json {
	j := Json{}

	// 实例化查询对象
	j.Query = query.NewQuery()

	// 实例化方法
	j.Dump = Dump
	j.Load = Load
	j.Dumps = Dumps
	j.Loads = Loads

	// 返回对象
	return &j
}

// Dumps 将Golang对象转换为json字符串
func Dumps(obj interface{}) (string, error) {
	rs, err := json.Marshal(obj)
	return string(rs), err
}

// Loads 将字符串转换为Golang对象
func Loads(str string, obj interface{}) error {
	err := json.Unmarshal([]byte(str), obj)
	return err
}

// Dump 将Golang对象写入到json文件
func Dump(filePath string, obj interface{}) error {

	// 创建json文件
	filePtr, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer filePtr.Close()

	// 写入json数据
	encoder := json.NewEncoder(filePtr)
	err = encoder.Encode(obj)
	return err
}

// Load 将json文件读取并转换为Golang对象
func Load(filePath string, obj interface{}) error {

	// 打开json文件
	filePtr, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer filePtr.Close()

	// 读取json文件
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(obj)
	return err
}
