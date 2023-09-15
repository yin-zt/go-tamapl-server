package logger

import (
	"github.com/cihub/seelog"
	"github.com/yin-zt/go-tamapl-server/pkg/config"
)

var (
	ServerLogger seelog.LoggerInterface
	CronLogger   seelog.LoggerInterface
	CmdLogger    seelog.LoggerInterface
	AccessLogger seelog.LoggerInterface
	resp         any
)

func Setup() {
	oneLogger, err := seelog.LoggerFromConfigAsBytes([]byte(config.LogCronConfigStr))
	if err != nil {
		resp = err
		panic(resp)
	}
	CronLogger = oneLogger

	twoLogger, err := seelog.LoggerFromConfigAsBytes([]byte(config.LogCmdConfigStr))
	if err != nil {
		resp = err
		panic(resp)
	}
	CmdLogger = twoLogger

	threeLogger, err := seelog.LoggerFromConfigAsBytes([]byte(config.LogServerConfigStr))
	if err != nil {
		resp = err
		panic(resp)
	}
	ServerLogger = threeLogger

	fourLogger, err := seelog.LoggerFromConfigAsBytes([]byte(config.LogAccessConfigStr))
	if err != nil {
		resp = err
		panic(resp)
	}
	AccessLogger = fourLogger
}
