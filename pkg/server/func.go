package server

import (
	"fmt"
	"github.com/coreos/etcd/client"
	"github.com/yin-zt/go-tamapl-server/pkg/config"
	"github.com/yin-zt/go-tamapl-server/pkg/utils/common"
	"github.com/yin-zt/go-tamapl-server/pkg/utils/logger"
	"os"
	"strings"
	"time"
)

func (ez *CliServer) Init(action string) {
	ParseConfig("")
	ez.InitComponent("init")
	ez.Etcd_host = Config().Etcd.Host
	etcdConf := &EtcdConf{User: Config().Etcd.User, Password: Config().Etcd.Pwd}
	ez.Util = &common.Common{}
	etcdStr := etcdConf.User + ":" + etcdConf.Password
	ez.Etcdbasicauth = "Basic" + ez.Util.Base64Encode(etcdStr)
	go Cli.InitEtcd()
	go Cli.InitUserAdmin()
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
	ez.Etcd_host = Config().Etcd.Host
	etcdconf := &EtcdConf{User: Config().Etcd.User, Password: Config().Etcd.Pwd}
	ez.Util = &common.Common{}
	str := etcdconf.User + ":" + etcdconf.Password
	ez.Etcdbasicauth = "Basic " + ez.Util.Base64Encode(str)

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

	if !ez.Util.IsExist(config.CONST_UPLOAD_DIR) {
		os.Mkdir(config.CONST_UPLOAD_DIR, 777)
	}

	// 检查程序中db、etcd、redis服务状态
	go func() {
		time.Sleep(time.Second * 2)

		status := Cli.checkstatus()

		logger.ServerLogger.Info(Cli.Util.JsonEncode(status))

		ticker := time.NewTicker(time.Minute)
		for {
			ez.ReportStatus()
			<-ticker.C
		}
	}()

}
