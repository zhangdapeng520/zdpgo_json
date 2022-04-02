package main

import (
	"github.com/zhangdapeng520/zdpgo_json/libs/gjson"
)

func main() {
	const json = `{"name": "Gilbert", "age": 61}
				  {"name": "Alexa", "age": 34}
				  {"name": "May", "age": 57}
				  {"name": "Deloise", "age": 44}`

	// 遍历每一行json
	gjson.ForEachLine(json, func(line gjson.Result) bool {
		println(line.String())
		return true
	})
}
