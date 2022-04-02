package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_json"
)

type account struct {
	Email    string  `json:"email"`
	password string  `json:"password"` // 不会处理私有变量
	Money    float64 `json:"money"`
}

type user struct {
	Name    string
	Age     int
	Roles   []string
	Skill   map[string]float64
	Account account
}

func main() {
	a := account{
		Email:    "张大鹏",
		password: "123456",
		Money:    100.5,
	}
	u := user{
		Name:    "张大鹏",
		Age:     27,
		Roles:   []string{"Owner", "Master"}, // 处理切片
		Account: a,
	}

	j := zdpgo_json.New()

	// 写入文件
	err := j.Dump("user.json", u)
	if err != nil {
		fmt.Println(err)
	}

	// 读取文件
	err = j.Load("user.json", &u)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(u)
}
