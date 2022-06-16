package main

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/zhangdapeng520/zdpgo_json"
)

func main() {
	/*
	   EndpointRule的内容大概如下
	   rule:
	     windows:
	       psh:
	         command: Get-WmiObject -class win32_operatingsystem | select -property *;

	     vul_tag: [["系统信息"],["win32_operatingsystem"]]
	*/
	data := make(map[string]interface{})
	windows := make(map[string]interface{})
	psh := make(map[string]string)
	psh["command"] = "Get-WmiObject -class win32_operatingsystem | select -property *;"
	windows["psh"] = psh
	data["windows"] = windows

	var vulTag [][]string
	vulTag = append(vulTag, []string{"系统信息"})
	vulTag = append(vulTag, []string{"win32_operatingsystem"})
	data["vul_tag"] = vulTag

	// 解析为JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Println("this json string : ", string(jsonData))

	// 重新解析为map
	var dataMap map[string]interface{}
	json.Unmarshal(jsonData, &dataMap)
	fmt.Println("read data map successful", dataMap)

	// 逐层提取
	for k, v := range dataMap {
		switch argValue := v.(type) {
		case map[string]interface{}: // get command
			fmt.Println("command", argValue)
			// 得到shell类型和命令
			for k1, v1 := range argValue {
				fmt.Println("===========", k, k1, v1)
				for k2, v2 := range v1.(map[string]interface{}) {
					fmt.Println("xxxxxxxxxxxxxxxx", k, k1, k2, v2)
				}
			}
		case []interface{}: // get vul tag
			var vulTag1 [][]string
			for _, v := range argValue {
				var t []string
				for _, vv := range v.([]interface{}) {
					t = append(t, vv.(string))
				}
				vulTag1 = append(vulTag1, t)
			}
			fmt.Println("vul tag", vulTag1)
		default:
			fmt.Println("unknown type")
		}
	}

	// 直接json提取
	jsonStr := `{"vul_tag":[["系统信息"],["win32_operatingsystem"]],"windows":{"psh":{"command":"Get-WmiObject -class win32_operatingsystem | select -property *;"}}}`
	j := zdpgo_json.New()
	vulTag2 := j.Query.Get(jsonStr, "vul_tag")
	fmt.Println("vul tag 2 = ", vulTag2, reflect.TypeOf(vulTag2))
	fmt.Println(vulTag2.Raw, reflect.TypeOf(vulTag2.Raw))
	var vulTag3 [][]string
	j.Loads(vulTag2.Raw, &vulTag3)
	fmt.Println("vul tag 3 = ", vulTag3, reflect.TypeOf(vulTag3), vulTag3[0][0])

	// 提取commoand
	command := j.Query.Get(jsonStr, "windows.psh.command")
	fmt.Println(command.Raw)
}
