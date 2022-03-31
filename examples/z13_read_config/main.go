package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_json"
)

type conf struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Db       string `json:"db"`
}

func main() {
	var c conf
	_ = zdpgo_json.ReadDefaultConfig(&c)
	fmt.Println("读取配置：", c.Username, c.Password, c.Host, c.Port, c.Db)
}
