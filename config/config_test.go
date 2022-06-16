package config

import (
	"fmt"
	"testing"
)

type conf struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Db       string `json:"db"`
}

// 测试读取配置
func TestJson_ReadConfig(t *testing.T) {
	var c conf
	// 读取公共配置
	_ = ReadConfig(&c, "config/config.json")
	fmt.Println(c, c.Username, c.Password)

	// 读取私密配置
	_ = ReadConfig(&c, "config/secret/.config.json")
	fmt.Println(c, c.Username, c.Password)

	// 同时读取多个
	var c1 conf
	_ = ReadConfig(&c1, "config/config.json", "config/secret/.config.json")
	fmt.Println(c1, c1.Username, c1.Password)
}

// 测试读取私密配置
func TestJson_ReadDefaultConfig(t *testing.T) {
	var c conf
	_ = ReadDefaultConfig(&c)
	fmt.Println(c, c.Username, c.Password, c.Host, c.Port, c.Db)
}
