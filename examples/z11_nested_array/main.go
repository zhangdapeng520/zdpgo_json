package main

import (
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

	// 获取每一行的lastName
	result := gjson.Get(json, "programmers.#.lastName")
	for _, name := range result.Array() {
		println(name.String())
	}

	// 查找lastName为Hunter的数据
	name := gjson.Get(json, `programmers.#(lastName="Hunter").firstName`)
	println(name.String())

	// 遍历数组
	result = gjson.Get(json, "programmers")
	result.ForEach(func(_, value gjson.Result) bool {
		println(value.String())
		return true // keep iterating
	})
}
