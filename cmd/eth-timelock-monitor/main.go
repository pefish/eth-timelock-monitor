package main

import (
	"github.com/pefish/eth-timelock-monitor/cmd/eth-timelock-monitor/command"
	"github.com/pefish/eth-timelock-monitor/version"
	"github.com/pefish/go-commander"
	go_logger "github.com/pefish/go-logger"
)

func main() {
	commanderInstance := commander.NewCommander(version.AppName, version.Version, version.AppName + " 是一个时间锁监控工具，祝你玩得开心。作者：pefish")
	commanderInstance.RegisterDefaultSubcommand(command.NewDefaultCommand())
	err := commanderInstance.Run()
	if err != nil {
		go_logger.Logger.Error(err)
	}
}
