package server

import (
	"fmt"
	"github.com/astaxie/beego/httplib"
	"github.com/garyburd/redigo/redis"
	"github.com/yin-zt/go-tamapl-server/pkg/utils/logger"
	"runtime/debug"
	"time"
)

// checkstatus检查redis\etcd\db的状态
func (ez *CliServer) checkstatus() map[string]string {
	data := make(map[string]string)
	data["redis"] = "fail"
	data["db"] = "fail"
	data["etcd"] = "fail"

	if ez.Util.CheckEnginer(Engine) {
		data["db"] = "ok"
	}

	if ez.Rp.ActiveCount() < ez.Rp.MaxActive {
		//		c := this.rp.Get()
		//		defer c.Close()
		//		c.Do("set", "hello", "world")
		ez.redisDo("set", "hello", "world")
		if result, err := redis.String(ez.redisDo("GET", "hello")); err == nil && result == "world" {
			data["redis"] = "ok"
		} else {
			fmt.Println(err)
		}
	}
	url := Config().EtcdGuest.Server[0] + Config().EtcdGuest.Prefix + "/hello"
	val, _ := ez.WriteEtcd(url, "world", "10")
	result := make(map[string]interface{})
	if err := json.Unmarshal([]byte(val), &result); err != nil {
		data["etcd"] = "fail"
	}

	if val, ok := result["node"]; ok {
		switch node := val.(type) {
		case map[string]interface{}:
			if k, o := node["value"]; o {
				if k == "world" {
					data["etcd"] = "ok"
				}
			}

		}
	}

	return data

}

// redisDo 用于连接redis池，并执行传入命令
func (ez *CliServer) redisDo(action string, args ...interface{}) (reply interface{}, err error) {
	c := ez.Rp.Get()
	defer c.Close()
	return c.Do(action, args...) //fuck
}

// WriteEtcd 往etcd指定的节点写入值，同时设置ttl时间
func (ez *CliServer) WriteEtcd(url string, value string, ttl string) (string, error) {

	req := httplib.Post(url)

	req.Header("Authorization", ez.Etcdbasicauth)
	req.Param("value", value)
	req.Param("ttl", ttl)
	req.SetTimeout(time.Second*10, time.Second*5)
	str, err := req.String()
	//	fmt.Println(str)
	if err != nil {
		logger.ServerLogger.Error(err)
		print(err)
	}
	return str, err

}

// ReportStatus
// todo: check db\redis\etcd status and report them to pl manager
func (ez *CliServer) ReportStatus() {
	var (
		resp any
	)
	defer func() {
		if re := recover(); re != resp {

			buffer := debug.Stack()
			logger.ServerLogger.Error(string(buffer))
		}
	}()

	data := ez.checkstatus()
	for k, v := range data {
		if v != "ok" && k != "db" {
			break
		}
	}
	fmt.Println("report pl status, attention!")

	logger.ServerLogger.Info(ez.Util.JsonEncode(data))

}
