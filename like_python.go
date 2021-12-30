package zdpgo_json

import (
	"os"

	jsoniter "github.com/json-iterator/go"
)


var json = jsoniter.ConfigCompatibleWithStandardLibrary

// 模拟Python的API

// 将Golang对象转换为json字符串
func Dumps(obj interface{}) (string, error) {
	rs, err := json.Marshal(obj)
	return string(rs), err
}

// 将字符串转换为Golang对象
func Loads(str string, obj interface{}) error {
	err := json.Unmarshal([]byte(str), obj)
	return err
}

// 将Golang对象写入到json文件
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

// 将json文件读取并转换为Golang对象
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
