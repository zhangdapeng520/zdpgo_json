# zdpgo_json
在Golang中便捷从处理Json文件，比如动态配置，增删改查

项目地址：https://github.com/zhangdapeng520/zdpgo_json

## 功能清单
- 与Python保持一致的API接口，dump对应Dump，load对应Load，dumps对应Dumps，loads对应Loads
- 支持json字符串的查询

## 版本历史
- 版本0.1.0 2022年2月16日 基本功能
- 版本0.1.1 2022年3月30日 读取配置
- 版本0.1.2 2022年4月2日 项目结构优化

## 使用示例

### 读写json文件
```go
package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_json"
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

func main() {
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

	j := zdpgo_json.New()

	// 写入文件
	err := j.Dump("user.json", u)
	if err != nil {
		fmt.Println(err)
	}

	// 读取文件
	err = j.Load("user.json", &u)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(u)
}
```

### 序列化和反序列化
```go
package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_json"
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

func main() {
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

	j := zdpgo_json.New()

	// 序列化
	jsonData, err := j.Dumps(u)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(jsonData)

	// 反序列化
	err = j.Loads(jsonData, &u)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(u)
}
```

### 从json字符串中查询数据
```go
package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_json"
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

func main() {
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

	j := zdpgo_json.New()

	// 序列化
	jsonData, err := j.Dumps(u)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(jsonData)

	// Get查询
	money := j.Query.Get(jsonData, "Account.money")
	fmt.Println("Get查询:", money, money.Float())
}
```
