package zdpgo_json

import (
	"fmt"
	"testing"
)

type account struct {
	Email    string  `json:"email"`
	password string  `json:"password"` // 不会处理私有变量
	Money    float64 `json:"money"`
}

type user struct {
	Name    string
	Age     int
	Roles   []string
	Skill   map[string]float64
	Account account
}

func getJson() *Json {
	return New()
}

// 测试写入和读取
func TestJson_DumpAndLoad(t *testing.T) {
	a := account{
		Email:    "张大鹏",
		password: "123456",
		Money:    100.5,
	}
	u := user{
		Name:    "张大鹏",
		Age:     27,
		Roles:   []string{"Owner", "Master"}, // 处理切片
		Account: a,
	}

	j := getJson()

	// 写入文件
	err := j.Dump("user.json", u)
	if err != nil {
		t.Error(err)
	}

	// 读取文件
	err = j.Load("user.json", &u)
	if err != nil {
		t.Error(err)
	}
	t.Log(u)
}

// 测试序列化和反序列化
func TestJson_DumpsAndLoads(t *testing.T) {
	a := account{
		Email:    "张大鹏",
		password: "123456",
		Money:    100.5,
	}
	u := user{
		Name:    "张大鹏",
		Age:     27,
		Roles:   []string{"Owner", "Master"}, // 处理切片
		Account: a,
	}

	j := getJson()

	// 序列化
	jsonData, err := j.Dumps(u)
	if err != nil {
		t.Error(err)
	}
	t.Log(jsonData)

	// 反序列化
	err = j.Loads(jsonData, &u)
	if err != nil {
		t.Error(err)
	}
	t.Log(u)
}

// 测试查询方法
func TestJson_query(t *testing.T) {
	a := account{
		Email: "张大鹏",
		Money: 100.5,
	}
	u := user{
		Name:    "张大鹏",
		Age:     27,
		Roles:   []string{"Owner", "Master"}, // 处理切片
		Account: a,
	}

	j := getJson()

	// 序列化
	jsonData, err := j.Dumps(u)
	fmt.Println(jsonData, err)

	// Get查询

}
