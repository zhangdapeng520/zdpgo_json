package config

import "github.com/zhangdapeng520/zdpgo_json/core/python"

// ReadConfig 读取配置，支持同时读取多个
func ReadConfig(configObj interface{}, configFileList ...string) error {
	for _, configFile := range configFileList {
		err := python.Load(configFile, configObj)
		if err != nil {
			return err
		}
	}
	return nil
}

// ReadDefaultConfig 读取默认配置。默认公共配置config/config.json，默认私密配置config/secret/.config.json
func ReadDefaultConfig(configObj interface{}) error {
	err := ReadConfig(configObj, "config/config.json", "config/secret/.config.json")
	return err
}
