package server

import (
	"fmt"
	"github.com/astaxie/beego/httplib"
	log "github.com/cihub/seelog"
	"github.com/garyburd/redigo/redis"
	"github.com/yin-zt/go-tamapl-server/pkg/utils/logger"
	"runtime/debug"
	"strings"
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
		logger.Error("check etcd server error: ", err)
	}

	logger.Info(result)
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
	logger.Info(ez.Etcdbasicauth)
	req.Header("Authorization", ez.Etcdbasicauth)
	req.Param("value", value)
	req.Param("ttl", ttl)
	req.SetTimeout(time.Second*10, time.Second*5)
	str, err := req.String()
	//	fmt.Println(str)
	if err != nil {
		logger.Error(err)
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
			logger.Error(string(buffer))
		}
	}()

	data := ez.checkstatus()
	for k, v := range data {
		if v != "ok" && k != "db" {
			break
		}
	}
	fmt.Println("report pl status, attention!")

	logger.Info(ez.Util.JsonEncode(data))

}

// InitEtcd 对etcd中的用户包含root和guest等用户设置auth/users和auth/roles读写权限
func (ez *CliServer) InitEtcd() {
	defer log.Flush()
	url := Config().Etcd.Host
	if strings.Index(url, "/v2") > 0 {
		url = url[0:strings.Index(url, "/v2")] + "/v2/"
	} else {
		url = url + "/v2/"
	}

	users := []string{Config().Etcd.User, Config().EtcdGuest.User}

	if !ez.Util.Contains("root", users) {
		msg := "etcd_root user must be root"
		log.Warn(msg)
		fmt.Println(msg)
		return
	}

	if Config().Etcd.Pwd == "" {
		msg := "etcd_root password must be not null"
		log.Warn(msg)
		fmt.Println(msg)
		return
	}

	for _, v := range users {
		data := "{\"role\":\"%s\",\"permissions\":{\"kv\":{\"read\":null,\"write\":null}}}"
		req := httplib.Put(url + fmt.Sprintf("auth/roles/%s", v))
		req.Body(fmt.Sprintf(data, v))
		log.Info(req.String())

	}
	for i, v := range users {
		data := "{\"user\":\"%s\",\"password\":\"%s\",\"roles\":[\"%s\"]}"
		req := httplib.Put(url + fmt.Sprintf("auth/users/%s", v))
		if i == 0 {
			req.Body(fmt.Sprintf(data, v, Config().Etcd.Pwd, v))
		} else {
			req.Body(fmt.Sprintf(data, v, Config().EtcdGuest.Password, v))
		}
		log.Info(req.String())
	}

	for i, v := range users {

		req := httplib.Put(url + fmt.Sprintf("auth/roles/%s", v))
		if i == 0 {
			data := "{\"role\":\"%s\",\"permissions\":{\"kv\":{\"read\":null,\"write\":null}},\"grant\":{\"kv\":{\"read\":[\"/keeper/*\"],\"write\":[\"/keeper/*\"]}}}"
			req.Body(fmt.Sprintf(data, v))
		} else {
			data := "{\"role\":\"%s\",\"permissions\":{\"kv\":{\"read\":null,\"write\":null}},\"grant\":{\"kv\":{\"read\":[\"/keeper/*\"],\"write\":null}}}"
			req.Body(fmt.Sprintf(data, v))
		}
		log.Info(req.String())
	}

	req := httplib.Put(url + "auth/enable")
	log.Info(req.String())
	log.Info("success to init etcd")

}

// InitUserAdmin 在TChUser表中插入一条admin的用户和密码数据
// 并且在TChAuth表中插入一条允许127.0.0.1以root登录的数据
func (ez *CliServer) InitUserAdmin() {
	userBean := new(TChUser)
	if ok, err := Engine.Where("Fuser=?", "admin").Get(userBean); err == nil && !ok {
		userBean.Fuser = "admin"
		userBean.Femail = "admin@web.com"
		userBean.Fpwd = ez.Util.MD5("admin")
		userBean.Fip = "127.0.0.1"
		userBean.Fstatus = 1
		if cnt, er := Engine.Insert(userBean); cnt > 0 {
			msg := "init admin user success,user=admin password=admin"
			log.Info(msg)
			fmt.Println(msg)
		} else {
			log.Error(er)
		}

	}
	auth := new(TChAuth)
	if ok, err := Engine.Where("Fsalt=?", "abc").Get(auth); err == nil && !ok {
		auth.Fuser = "root"
		auth.Fsudo = 1
		auth.Fenable = 1
		auth.FsudoIps = "*"
		auth.Fdesc = "local test"
		auth.Ftoken = "abc"
		auth.Fip = "127.0.0.1"
		if cnt, er := Engine.Insert(auth); cnt > 0 {
			msg := "init auth token abc success,ip=127.0.0.1 token=abc sudo=1"
			log.Info(msg)
			fmt.Println(msg)
		} else {
			log.Error(er)
		}

	}

}
