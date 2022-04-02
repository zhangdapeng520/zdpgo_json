package main

import (
	"encoding/json"
	"fmt"
	"github.com/zhangdapeng520/zdpgo_json/core/python"
	"log"
)

type Account struct {
	Email    string  `json:"email"`
	password string  `json:"password"` // 不会处理私有变量
	Money    float64 `json:"money"`
}

// GoMarshal 使用Golang的Marshal方法
func GoMarshal() {
	account := Account{
		Email:    "张大鹏",
		password: "123456",
		Money:    100.5,
	}

	rs, err := json.Marshal(account)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(rs)
	fmt.Println(string(rs))
}

// PythonDumps 使用类似Python的dumps方法
func PythonDumps() {
	account := Account{
		Email:    "张大鹏",
		password: "123456",
		Money:    100.5,
	}
	result, _ := python.Dumps(account)
	fmt.Println(result)
}

func GoUnmarshal() {
	str := "{\"access_token\":\"uAUS6o5g-9rFWjYt39LYa7TKqiMVsIfCGPEN4IZzdAk5-T-ryVhL7xb8kYciuU_m\",\"expires_in\":7200}"
	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(str), &dat); err == nil {
		fmt.Println(dat)
		fmt.Println(dat["expires_in"])
	} else {
		fmt.Println(err)
	}
}

func PythonLoads() {
	str := "{\"access_token\":\"uAUS6o5g-9rFWjYt39LYa7TKqiMVsIfCGPEN4IZzdAk5-T-ryVhL7xb8kYciuU_m\",\"expires_in\":7200}"
	var dat map[string]interface{}
	_ = python.Loads(str, &dat)
	fmt.Println(dat)

	s := `{"email":"张大鹏","money":100.5}`
	var account Account
	_ = python.Loads(s, &account)
	fmt.Println(account, account.Email, account.Money)
}

func main() {
	GoMarshal()
	PythonDumps()
	GoUnmarshal()
	PythonLoads()
}
