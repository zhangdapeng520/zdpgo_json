package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_json/libs/gjson"
)

func main() {
	const json = `{
					"programmers": [
						{
						"firstName": "Janet", 
						"lastName": "McLaughlin", 
						}, {
						"firstName": "Elliotte", 
						"lastName": "Hunter", 
						}, {
						"firstName": "Jason", 
						"lastName": "Harold", 
						}
					]
				}`

	// 判断是否为json字符串
	if !gjson.Valid(json) {
		fmt.Println("json数据格式校验失败")
	}

	// 判断是否存在数据
	value := gjson.Get(json, "name.last")
	if !value.Exists() {
		println("no last name")
	} else {
		println(value.String())
	}

	// Or as one step
	if gjson.Get(json, "name.last").Exists() {
		println("has a last name")
	}
}
