package command

import (
	"flag"

	api_strategy "github.com/pefish/go-core-strategy/api-strategy"
	"github.com/pefish/remote-http-command/pkg/constant"
	"github.com/pefish/remote-http-command/pkg/global"
	"github.com/pefish/remote-http-command/pkg/route"
	"github.com/pefish/remote-http-command/version"

	"github.com/pefish/go-commander"
	go_config "github.com/pefish/go-config"
	"github.com/pefish/go-core/driver/logger"
	global_api_strategy "github.com/pefish/go-core/global-api-strategy"
	"github.com/pefish/go-core/service"
	go_logger "github.com/pefish/go-logger"
	task_driver "github.com/pefish/go-task-driver"
)

type DefaultCommand struct {
}

func NewDefaultCommand() *DefaultCommand {
	return &DefaultCommand{}
}

func (dc *DefaultCommand) DecorateFlagSet(flagSet *flag.FlagSet) error {
	flagSet.String("server-host", "0.0.0.0", "The host of web server.")
	flagSet.Int("server-port", 8000, "The port of web server.")
	flagSet.String("username", "username", "The username of basic auth.")
	flagSet.String("password", "password", "The password of basic auth.")
	return nil
}

func (dc *DefaultCommand) OnExited(command *commander.Commander) error {
	return nil
}

func (dc *DefaultCommand) Init(command *commander.Commander) error {
	service.Service.SetName(version.AppName)
	logger.LoggerDriverInstance.Register(go_logger.Logger)

	err := go_config.ConfigManagerInstance.Unmarshal(&global.GlobalConfig)
	if err != nil {
		return err
	}
	service.Service.SetHost(global.GlobalConfig.ServerHost)
	service.Service.SetPort(global.GlobalConfig.ServerPort)
	service.Service.SetPath(`/api`)
	global_api_strategy.ParamValidateStrategyInstance.SetErrorCode(constant.PARAM_ERROR)

	global.BasicAuthStrategy.SetParams(api_strategy.BasicAuthParams{
		Username: global.GlobalConfig.Username,
		Password: global.GlobalConfig.Password,
	})

	service.Service.SetRoutes(route.CmdRoute)

	return nil
}

func (dc *DefaultCommand) Start(command *commander.Commander) error {

	taskDriver := task_driver.NewTaskDriver()
	taskDriver.Register(service.Service)

	taskDriver.RunWait(command.Ctx)

	return nil
}
