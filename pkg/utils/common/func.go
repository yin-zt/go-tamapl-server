package common

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"net"
	"strings"
	"time"
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
