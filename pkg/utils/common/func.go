package common

import (
	"encoding/base64"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"net"
	"os"
	"strings"
)

// GetPulicIP 作用是获取本地IP地址，且需要能与外部DNS 8.8.8.8:80 实现udp通信的
func (this *Common) GetPulicIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "127.0.0.1"
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().String()
	idx := strings.LastIndex(localAddr, ":")
	return localAddr[0:idx]
}

// Base64Encode 对输入字符串基于base64进行编码
func (this *Common) Base64Encode(str string) string {

	return base64.StdEncoding.EncodeToString([]byte(str))
}

// IsExist 用于判断文件是否存在
func (this *Common) IsExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

// JsonEncode 结构体入参转为字符串
func (this *Common) JsonEncode(v interface{}) string {

	if v == nil {
		return ""
	}
	vv := map[string]string{"db": "ok", "etcd": "ok", "redis": "ok"}
	tools := jsoniter.ConfigCompatibleWithStandardLibrary
	data, _ := tools.Marshal(vv)
	fmt.Println(data)
	jbyte, err := json.Marshal(v)
	fmt.Println(jbyte)
	if err == nil {
		return string(jbyte)
	} else {
		return ""
	}

}
