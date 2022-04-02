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

	// 序列化
	jsonData, err := j.Dumps(u)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(jsonData)

	// Get查询
	money := j.Query.Get(jsonData, "Account.money")
	fmt.Println("Get查询:", money, money.Float())
}
