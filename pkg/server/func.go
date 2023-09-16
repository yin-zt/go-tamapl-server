package server

import (
	"fmt"
	"github.com/coreos/etcd/client"
	"github.com/yin-zt/go-tamapl-server/pkg/config"
	"github.com/yin-zt/go-tamapl-server/pkg/utils/logger"
	"strings"
	"time"
)

func (ez *CliServer) Init(action string) {
	ParseConfig("")
	ez.InitComponent("init")
}

func (ez *CliServer) InitComponent(action string) {
	cfg := client.Config{
		Endpoints:               []string{Config().Etcd.Host},
		Transport:               client.DefaultTransport,
		Username:                Config().Etcd.User,
		Password:                Config().Etcd.Pwd,
		HeaderTimeoutPerRequest: time.Second,
	}

	for i, v := range Config().EtcdGuest.Server {
		if strings.Index(v, "/v2/") <= 0 {
			Config().EtcdGuest.Server[i] = Config().EtcdGuest.Server[i] + config.CONST_ETCD_PREFIX
		}
	}

	c, err := client.New(cfg)
	if err != nil {
		logger.ServerLogger.Error(err)
	}
	Cli.Kapi = client.NewKeysAPI(c)
	Cli.EtcdClent = c

	DbConfigMap := map[string]string{
		"dbtype":   Config().Db.Type,
		"user":     Config().Db.User,
		"password": Config().Db.Password,
		"host":     Config().Db.Host,
		"port":     Config().Db.Port,
		"db":       Config().Db.Db,
	}
	createdEnginer, err := Cli.Util.InitEngine(DbConfigMap)
	if err != nil {
		logger.ServerLogger.Error("init db engine occur error!")
		resp = "error occur, run !!!"
		panic(resp)
	}
	Engine = createdEnginer
	fmt.Println(createdEnginer)

	if config.AutoCreatTable {
		// 在数据库上创建相应的表
		if err := Engine.Sync2(new(TChUser), new(TChResults), new(TChHeartbeat),
			new(TChLog), new(TChResultsHistory), new(TChConfig)); err != nil {
			logger.ServerLogger.Error(err.Error())
		}
	}

	RedisConfigMap := map[string]interface{}{
		"maxIdle":        Config().Redis.MaxIdle,
		"maxActive":      Config().Redis.MaxActive,
		"idleTimeout":    Config().Redis.IdleTimeout,
		"address":        Config().Redis.Address,
		"pwd":            Config().Redis.Pwd,
		"db":             Config().Redis.DB,
		"connectTimeout": Config().Redis.ConnectTimeout,
	}
	redisPool, err := Cli.Util.InitRedisPool(RedisConfigMap)
	if err != nil {
		logger.ServerLogger.Error("init redis pool error, go away")
		resp = "fail to init redis pool"
		panic(resp)
	}
	Cli.Rp = redisPool
}