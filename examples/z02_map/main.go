package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type User struct {
	Name  string
	Age   int
	Roles []string
	Skill map[string]float64
}

func main() {

	// 处理map
	skill := make(map[string]float64)
	skill["python"] = 99.5
	skill["elixir"] = 90
	skill["ruby"] = 80.0

	user := User{
		Name:  "张大鹏",
		Age:   27,
		Roles: []string{"Owner", "Master"}, // 处理切片
		Skill: skill,
	}

	rs, err := json.Marshal(user)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(rs))
}
