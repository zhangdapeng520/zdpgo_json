package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_json/libs/gjson"
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

	// 查找字符串
	value := gjson.Get(json, "name.last")
	println(value.String())

	// 获取数组长度
	arrLen := gjson.Get(json, "children.#")
	println(arrLen.Int())

	// 获取数组指定索引元素
	arrIndex := gjson.Get(json, "children.1")
	println(arrIndex.String())

	// 模糊匹配获取数组指定索引元素
	arrLikeIndex := gjson.Get(json, "child*.2")
	println(arrLikeIndex.String())

	// 模糊匹配获取数组指定索引元素
	arrLikeOneIndex := gjson.Get(json, "c?ildren.0")
	println(arrLikeOneIndex.String())

	// 键本身包含小数点，使用转义字符
	trans := gjson.Get(json, `fav\.movie`) // 注意：不要用双引号
	println(trans.String())

	// 取所有数组的指定元素
	arrAllFrist := gjson.Get(json, "friends.#.first")
	println(arrAllFrist.Array())
	for _, v := range arrAllFrist.Array() {
		fmt.Print(v, " ")
	}
	fmt.Println()

	// 取指定数组的指定元素
	arrFirstFrist := gjson.Get(json, "friends.1.first")
	println(arrFirstFrist.String())
}
