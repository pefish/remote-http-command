package main

import (
	"github.com/pefish/go-commander"
	go_logger "github.com/pefish/go-logger"
	"github.com/pefish/remote-http-command/cmd/remote-http-command/command"
	"github.com/pefish/remote-http-command/version"
)

func main() {
	commanderInstance := commander.NewCommander(
		version.AppName,
		version.Version,
		version.AppName+" 是一个使用 HTTP 请求执行远程命令的工具。作者：pefish",
	)
	commanderInstance.RegisterDefaultSubcommand(&commander.SubcommandInfo{
		Desc:       "",
		Args:       nil,
		Subcommand: command.NewDefaultCommand(),
	})
	err := commanderInstance.Run()
	if err != nil {
		go_logger.Logger.Error(err)
	}
}
