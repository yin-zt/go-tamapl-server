package server

import (
	"fmt"
	"github.com/yin-zt/go-tamapl-server/pkg/config"
	"github.com/yin-zt/go-tamapl-server/pkg/utils/logger"
	"io/ioutil"
	"os"
	"strings"
	"sync/atomic"
	"unsafe"
)

func Config() *ServerConfig {
	return (*ServerConfig)(atomic.LoadPointer(&ptr))
}

func ParseConfig(filePath string) {
	var (
		data []byte
		resp any
	)

	if filePath == "" {
		data = []byte(strings.TrimSpace(config.CfgJson))
	} else {
		file, err := os.Open(filePath)
		if err != nil {
			resp = fmt.Sprintln("open file path:", filePath, "error:", err)
			panic(resp)
		}

		defer file.Close()

		FileName = filePath

		data, err = ioutil.ReadAll(file)
		if err != nil {
			resp = fmt.Sprintln("file path:", filePath, " read all error:", err)
			panic(resp)
		}
	}

	// 将配置文件内容解析出来，并存入c变量中
	var c ServerConfig

	fmt.Println(string(data))
	if err := json.Unmarshal(data, &c); err != nil {
		fmt.Println(err)
		resp = fmt.Sprintln("file path:", filePath, "json unmarshal error:", err)
		panic(resp)
	}

	//// 根据配置文件中数据库的类型进行逻辑处理，db默认为sqlite3
	//// 如果c.Db.Type为空，则再次使用本地ip渲染配置文件
	//if c.Db.Type == "" {
	//	ip := cli.Util.GetPulicIP()
	//	cfg := fmt.Sprintf(cfgJson, "http://"+ip+":9160", cli.Util.GetUUID(), ip, cli.Util.GetUUID(), ip)
	//	cli.Util.WriteFile(CONST_CFG_FILE_NAME, string(cfg))
	//	if !cli.Util.IsExist(CONST_LOCAL_CFG_FILE_NAME) {
	//		cli.Util.WriteFile(CONST_LOCAL_CFG_FILE_NAME, string(cfg))
	//	}
	//	if err := json.Unmarshal([]byte(cfg), &c); err != nil {
	//		resp = fmt.Sprintln("file path:", filePath, "json unmarshal error:", err)
	//		panic(resp)
	//	}
	//}

	atomic.StorePointer(&ptr, unsafe.Pointer(&c))
	logger.Info("config parse success")
}
