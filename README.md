# zdpgo_json
在Golang中便捷从处理Json文件，比如动态配置，增删改查

项目地址：https://github.com/zhangdapeng520/zdpgo_json

## 功能清单
- 与Python保持一致的API接口，dump对应Dump，load对应Load，dumps对应Dumps，loads对应Loads
- 采用json-iterator进行json的编码和解码，效率更高
- 整合gjson，支持json字符串的查询

## 版本历史
- 版本0.1.0 2022年2月16日 基本功能

## 一、快速入门

### 1.1 读写json字符串
```go
package main

import (
	`fmt`
	
	`github.com/zhangdapeng520/zdpgo_json`
)

type Account struct {
	Email    string  `json:"email"`
	password string  `json:"password"` // 不会处理私有变量
	Money    float64 `json:"money"`
}

// 使用类似Python的dumps方法
func PythonDumps() {
	account := Account{
		Email:    "张大鹏",
		password: "123456",
		Money:    100.5,
	}
	result, _ := zdpgo_json.Dumps(account)
	fmt.Println(result)
}

func PythonLoads() {
	str := "{\"access_token\":\"uAUS6o5g-9rFWjYt39LYa7TKqiMVsIfCGPEN4IZzdAk5-T-ryVhL7xb8kYciuU_m\",\"expires_in\":7200}"
	var dat map[string]interface{}
	_ = zdpgo_json.Loads(str, &dat)
	fmt.Println(dat)
	
	s := `{"email":"张大鹏","money":100.5}`
	var account Account
	_ = zdpgo_json.Loads(s, &account)
	fmt.Println(account, account.Email, account.Money)
}

func main() {
	PythonDumps()
	PythonLoads()
}
```

### 1.2 写入json文件
```go
package main

import (
	`github.com/zhangdapeng520/zdpgo_json`
)

type Account struct {
	Email    string  `json:"email"`
	password string  `json:"password"` // 不会处理私有变量
	Money    float64 `json:"money"`
}

type User struct {
	Name    string
	Age     int
	Roles   []string
	Skill   map[string]float64
	Account Account
}

func PythonDump() {
	account := Account{
		Email:    "张大鹏",
		password: "123456",
		Money:    100.5,
	}
	
	// 处理map
	skill := make(map[string]float64)
	skill["python"] = 99.5
	skill["elixir"] = 90
	skill["ruby"] = 80.0
	
	user := User{
		Name:    "rsj217",
		Age:     27,
		Roles:   []string{"Owner", "Master"},
		Skill:   skill,
		Account: account,
	}
	
	// 写json文件
	// 创建文件
	filePath := "./examples/z04_write_json/user1.json"
	_ = zdpgo_json.Dump(filePath, user)
}
func main() {
	PythonDump()
}
```

### 1.3读取json文件
```go
package main

import (
	`fmt`
	`github.com/zhangdapeng520/zdpgo_json`
)

type Account struct {
	Email    string  `json:"email"`
	password string  `json:"password"` // 不会处理私有变量
	Money    float64 `json:"money"`
}

type User struct {
	Name    string
	Age     int
	Roles   []string
	Skill   map[string]float64
	Account Account
}

func PythonLoad() {
	filePath := "./examples/z05_read_json/user.json"
	var user User
	_ = zdpgo_json.Load(filePath, &user)
	fmt.Println(user)
}

func main() {
	PythonLoad()
}
```

## 二、高级功能

### 2.1 查询json字符串
```go
package main

import (
	"fmt"

	"github.com/zhangdapeng520/zdpgo_json"
)


func main() {
	const json = `{"name":{"first":"dapeng","last":"zhang"},"age":47, "gender":true}`

	// 查找字符串
	value := zdpgo_json.Get(json, "name.last")
	println(value.String())

	// 查找数字
	age := zdpgo_json.Get(json, "age")
	fmt.Println(age.Int())

	// 查找布尔值
	gender := zdpgo_json.Get(json, "gender")
	fmt.Println(gender.Bool())
}
```
