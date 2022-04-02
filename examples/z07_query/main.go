package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_json/libs/gjson"
)

func main() {
	const json = `{"name":{"first":"dapeng","last":"zhang"},"age":47, "gender":true}`

	// 查找字符串
	value := gjson.Get(json, "name.last")
	println(value.String())

	// 查找数字
	age := gjson.Get(json, "age")
	fmt.Println(age.Int())

	// 查找布尔值
	gender := gjson.Get(json, "gender")
	fmt.Println(gender.Bool())
}
