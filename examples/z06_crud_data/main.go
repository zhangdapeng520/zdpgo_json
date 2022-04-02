package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_json/core/python"
	"sync"
)

// ConfigST 配置对象
type ConfigST struct {
	mutex     sync.RWMutex
	Server    ServerST            `json:"server"`  // 服务配置
	Streams   map[string]StreamST `json:"streams"` // 流配置
	LastError error               // 最后的错误
}

// ServerST 服务的相关配置
type ServerST struct {
	HTTPPort      string   `json:"http_port"`       // 端口号
	ICEServers    []string `json:"ice_servers"`     // ice服务列表
	ICEUsername   string   `json:"ice_username"`    // ice用户名
	ICECredential string   `json:"ice_credential"`  // ice证书
	WebRTCPortMin uint16   `json:"webrtc_port_min"` // webtrc最小端口号
	WebRTCPortMax uint16   `json:"webrtc_port_max"` // webrtc最大端口号
}

// StreamST 流媒体相关配置
type StreamST struct {
	URL          string            `json:"url"`    // 路径
	Status       bool              `json:"status"` // 状态
	OnDemand     bool              `json:"on_demand"`
	DisableAudio bool              `json:"disable_audio"` // 禁用音频
	Debug        bool              `json:"debug"`         // debug模式
	RunLock      bool              `json:"-"`             // 运行锁
	Codecs       []string          // 编码
	Cl           map[string]viewer // 连接
}

type viewer struct {
	c chan string
}

func main() {
	filePath := "examples\\z06_crud_data\\config.json"
	var data ConfigST
	_ = python.Load(filePath, &data)
	fmt.Println(data.Streams)

	// 添加
	newStream := StreamST{
		URL:          "rtsp://admin:xx123456@111.198.61.222:9999/h264/ch1/main/av_stream",
		DisableAudio: true,
	}
	data.Streams["test1"] = newStream

	// 重新写入
	_ = python.Dump(filePath, data)

	// 删除
	delete(data.Streams, "test1")
	// 重新写入
	_ = python.Dump(filePath, data)
}
