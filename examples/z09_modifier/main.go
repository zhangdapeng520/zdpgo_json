package main

import (
	"fmt"
	"strings"

	"github.com/zhangdapeng520/zdpgo_json"
)


func main() {
	const json = `{
					"name": {"first": "Tom", "last": "Anderson"},
					"age":37,
					"children": ["Sara","Alex","Jack"],
					"fav.movie": "Deer Hunter",
					"friends": [
						{"first": "Dale", "last": "Murphy", "age": 44, "nets": ["ig", "fb", "tw"]},
						{"first": "Roger", "last": "Craig", "age": 68, "nets": ["fb", "tw"]},
						{"first": "Jane", "last": "Murphy", "age": 47, "nets": ["ig", "tw"]}
					]
				}`
	
	/*
	@reverse: Reverse an array or the members of an object.
	@ugly: Remove all whitespace from a json document.
	@pretty: Make the json document more human readable.
	@this: Returns the current element. It can be used to retrieve the root element.
	@valid: Ensure the json document is valid.
	@flatten: Flattens an array.
	@join: Joins multiple objects into a single object.
	@keys: Returns an array of keys for an object.
	@values: Returns an array of values for an object.
	*/

	// 自定义过滤器
	zdpgo_json.AddModifier("case", func(json, arg string) string {
		if arg == "upper" {
			return strings.ToUpper(json)
		}
		if arg == "lower" {
			return strings.ToLower(json)
		}
		return json
	})

	// 使用过滤器
	value := zdpgo_json.Get(json, "children|@case:upper")
	for _, v:= range value.Array(){
		fmt.Print(v, " ")
	}
	fmt.Println()

	value = zdpgo_json.Get(json, "children|@case:lower")
	for _, v:= range value.Array(){
		fmt.Print(v, " ")
	}
	fmt.Println()
}