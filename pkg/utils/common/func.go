package common

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"net"
	"os"
	"reflect"
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

// Contains 方法的作用是判断obj是否被包含在arrayobj中，第二个参数是数组|集合
func (this *Common) Contains(obj interface{}, arrayobj interface{}) bool {
	targetValue := reflect.ValueOf(arrayobj)
	switch reflect.TypeOf(arrayobj).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return true
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true
		}
	}
	return false
}

func (this *Common) MD5(str string) string {

	md := md5.New()
	md.Write([]byte(str))
	return fmt.Sprintf("%x", md.Sum(nil))
}
