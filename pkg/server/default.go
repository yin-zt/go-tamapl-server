package server

import (
	"github.com/coreos/etcd/client"
	"github.com/garyburd/redigo/redis"
	"github.com/yin-zt/go-tamapl-server/pkg/utils/common"
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
