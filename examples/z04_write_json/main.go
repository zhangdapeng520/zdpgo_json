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
	account := Account{
		Email:    "张大鹏",
		password: "123456",
		Money:    100.5,
	}

	// 处理map
	skill := make(map[string]float64)
	skill["python"] = 99.5
	skill["elixir"] = 90
	skill["ruby"] = 80.0

	user := User{
		Name:    "rsj217",
		Age:     27,
		Roles:   []string{"Owner", "Master"},
		Skill:   skill,
		Account: account,
	}

	// 写json文件
	// 创建文件
	filePtr, err := os.Create("./examples/z04_write_json/user.json")
	if err != nil {
		fmt.Println("文件创建失败", err.Error())
		return
	}
	defer filePtr.Close()

	// 创建Json编码器
	encoder := json.NewEncoder(filePtr)
	err = encoder.Encode(user)
	if err != nil {
		fmt.Println("编码错误", err.Error())
	} else {
		fmt.Println("编码成功")
	}
}

func PythonDump() {
	account := Account{
		Email:    "张大鹏",
		password: "123456",
		Money:    100.5,
	}

	// 处理map
	skill := make(map[string]float64)
	skill["python"] = 99.5
	skill["elixir"] = 90
	skill["ruby"] = 80.0

	user := User{
		Name:    "rsj217",
		Age:     27,
		Roles:   []string{"Owner", "Master"},
		Skill:   skill,
		Account: account,
	}

	// 写json文件
	// 创建文件
	filePath := "./examples/z04_write_json/user1.json"
	_ = zdpgo_json.Dump(filePath, user)
}
func main() {
	GoMarshal()
	PythonDump()
}
