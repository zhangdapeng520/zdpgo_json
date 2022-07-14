package main

import (
	"fmt"

	"github.com/zhangdapeng520/zdpgo_json/jsoniter"
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
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

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

	// 序列化
	jsonBytes, err := json.Marshal(u)
	if err != nil {
		panic(err)
	}

	// 反序列化
	var u1 user
	err = json.Unmarshal(jsonBytes, &u1)
	if err != nil {
		panic(err)
	}

	fmt.Println("成功：", u1, u1.Name, u1.Age, u1.Roles, u1.Account)
}
