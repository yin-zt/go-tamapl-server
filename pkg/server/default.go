package server

import (
	"github.com/coreos/etcd/client"
	"github.com/garyburd/redigo/redis"
	"github.com/go-xorm/xorm"

	jsoniter "github.com/json-iterator/go"
	"github.com/yin-zt/go-tamapl-server/pkg/utils/common"
	"unsafe"
)

var (
	Cli      = &CliServer{Util: &common.Common{}}
	ptr      unsafe.Pointer
	FileName string
	json     = jsoniter.ConfigCompatibleWithStandardLibrary
	Engine   *xorm.Engine
	resp     any
)

type CliServer struct {
	Etcd_host     string
	Etcdbasicauth string
	Util          *common.Common
	Rp            *redis.Pool
	EtcdClent     client.Client  // etcd client实例
	Kapi          client.KeysAPI // etcd 交互接口调用
	EtcdDelKeys   chan string
}

type Etcd struct {
	Host string `json:"host"`
	User string `json:"user"`
	Pwd  string `json:"password"`
}

type HeartBeatEtcd struct {
	Prefix   string   `json:"prefix"`
	User     string   `json:"user"`
	Password string   `json:"password"`
	Server   []string `json:"server"`
}

type DB struct {
	Type     string `json:"type"`
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Db       string `json:"database"`
}

type Redis struct {
	Address        string `json:"address"`
	Pwd            string `json:"pwd"`
	MaxIdle        int    `json:"maxIdle"`
	MaxActive      int    `json:"maxActive"`
	IdleTimeout    int    `json:"idleTimeout"`
	ConnectTimeout int    `json:"connectTimeout"`
	DB             int    `json:"db"`
}

type ServerConfig struct {
	Etcd            Etcd          `json:"etcd_root"`
	EtcdGuest       HeartBeatEtcd `json:"etcd"`
	Db              DB            `json:"db"`
	Redis           Redis         `json:"redis"`
	EtcdValueExpire int           `json:"etcd_value_expire"`
}
