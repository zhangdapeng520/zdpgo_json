package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/zhangdapeng520/zdpgo_json"
)

type Account struct {
	Email    string  `json:"email"`
	password string  `json:"password"` // 不会处理私有变量
	Money    float64 `json:"money"`
}

type User struct {
	Name    string
	Age     int
	Roles   []string
	Skill   map[string]float64
	Account Account
}

func GoMarshal() {
	filePtr, err := os.Open("./examples/z05_read_json/user.json")
	if err != nil {
		fmt.Println("文件打开失败 [Err:%s]", err.Error())
		return
	}
	defer filePtr.Close()

	var user User
	// 创建json解码器
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(&user)
	if err != nil {
		fmt.Println("解码失败", err.Error())
	} else {
		fmt.Println("解码成功")
		fmt.Println(user)
	}
}

func PythonLoad() {
	filePath := "./examples/z05_read_json/user.json"
	var user User
	_ = zdpgo_json.Load(filePath, &user)
	fmt.Println(user)
}

func main() {
	GoMarshal()
	PythonLoad()
}
