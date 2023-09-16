package common

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/spf13/cast"
	"github.com/yin-zt/go-tamapl-server/pkg/utils/logger"
	"time"
)

// InitEngine 根据传入的配置信息初始化数据库驱动
func (this *Common) InitEngine(c map[string]string) (*xorm.Engine, error) {

	url := "%s:%s@tcp(%s:%s)/%s?charset=utf8"
	url = fmt.Sprintf(url, c["user"], c["password"], c["host"], c["port"], c["db"])
	dbtype := c["dbtype"]

	fmt.Println(url)
	fmt.Println(dbtype)
	//if Config().Debug {
	//	fmt.Println(url)
	//	log.Info(url)
	//}

	_enginer, er := xorm.NewEngine(dbtype, url)
	if er == nil /*&& this.CheckEnginer(_enginer)*/ {
		_enginer.SetConnMaxLifetime(time.Duration(60) * time.Second)
		_enginer.SetMaxIdleConns(0)
		//		_enginer.ShowSQL(true)
		return _enginer, nil
	} else {
		return nil, er
	}

}

// 根据配置文件的内容初始化redis pool
func (this *Common) InitRedisPool(c map[string]interface{}) (*redis.Pool, error) {

	pool := &redis.Pool{
		MaxIdle:     cast.ToInt(c["maxIdle"]),
		MaxActive:   cast.ToInt(c["maxActive"]),
		IdleTimeout: time.Duration(cast.ToInt(c["idleTimeout"])) * time.Second,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", cast.ToString(c["address"]),
				redis.DialConnectTimeout(time.Duration(cast.ToInt(c["connectTimeout"]))*time.Second),
				redis.DialPassword(cast.ToString(c["pwd"])),
				redis.DialDatabase(cast.ToInt(c["db"])),
			)
			if err != nil {
				fmt.Println(err)
				logger.ServerLogger.Error(err)
			}
			return conn, err

		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("ping")
			if err != nil {
				logger.ServerLogger.Error(err)
				return err
			}
			return err
		},
	}
	return pool, nil
}

// CheckEnginer 用于检查数据库驱动是否正常，通过执行 select 1的方式
func (this *Common) CheckEnginer(engine *xorm.Engine) bool {
	ret := false
	s := "select 1"
	if rows, err := engine.Query(s); err == nil {
		if len(rows) > 0 {
			if v, ok := rows[0]["1"]; ok {
				if string(v) == "1" {
					return true
				} else {
					return false
				}

			}
		}
	} else {
		fmt.Println(err)
	}

	return ret

}
