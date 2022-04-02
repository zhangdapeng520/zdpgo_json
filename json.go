package zdpgo_json

import (
	"github.com/zhangdapeng520/zdpgo_json/core/python"
	"github.com/zhangdapeng520/zdpgo_json/core/query"
)

// Json 处理json的核心对象
type Json struct {
	Query *query.Query // 查询核心对象

	// 方法列表
	Dump  func(filePath string, obj interface{}) error
	Load  func(filePath string, obj interface{}) error
	Dumps func(obj interface{}) (string, error)
	Loads func(str string, obj interface{}) error
}

// New 创建新的处理json的对象示例
func New() *Json {
	j := Json{}

	// 实例化查询对象
	j.Query = query.NewQuery()

	// 实例化方法
	j.Dump = python.Dump
	j.Load = python.Load
	j.Dumps = python.Dumps
	j.Loads = python.Loads

	// 返回对象
	return &j
}
