package main

import (
	"github.com/zhangdapeng520/zdpgo_json"
)


func main() {
	const json = `{"name": "Gilbert", "age": 61}
				  {"name": "Alexa", "age": 34}
				  {"name": "May", "age": 57}
				  {"name": "Deloise", "age": 44}`
	
	// 遍历每一行json
	zdpgo_json.ForEachLine(json, func(line zdpgo_json.Result) bool{
		println(line.String())
		return true
	})
}