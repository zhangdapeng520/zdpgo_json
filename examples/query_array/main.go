package main

import (
	"fmt"

	"github.com/zhangdapeng520/zdpgo_json"
)

func main() {
	jsonStr := `[{"response_status":true}]`
	j := zdpgo_json.New()

	// 查询数组
	result := j.Query.Get(jsonStr, "[0]")
	fmt.Println(result)
	fmt.Println(result.Exists())

	// 查询数组
	result = j.Query.Get(jsonStr, "[0].0")
	fmt.Println(result.Map()["response_status"])

	// 完整写法
	result = j.Query.Get(jsonStr, "[0].0")
	if result.Exists() {
		response_status := result.Get("response_status")
		if response_status.Exists() {
			fmt.Println(response_status.Bool())
		}
	}
}
