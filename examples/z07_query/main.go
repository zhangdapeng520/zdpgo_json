package main

import (
	"fmt"

	"github.com/zhangdapeng520/zdpgo_json"
)

const json = `{"name":{"first":"dapeng","last":"zhang"},"age":47, "gender":true}`

func main() {
	// 查找字符串
	value := zdpgo_json.Get(json, "name.last")
	println(value.String())

	// 查找数字
	age := zdpgo_json.Get(json, "age")
	fmt.Println(age.Int())

	// 查找布尔值
	gender := zdpgo_json.Get(json, "gender")
	fmt.Println(gender.Bool())
}