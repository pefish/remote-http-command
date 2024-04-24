package route

import (
	"github.com/pefish/remote-http-command/pkg/controller"
	"github.com/pefish/remote-http-command/pkg/global"

	"github.com/pefish/go-core/api"
	"github.com/pefish/go-http/gorequest"
)

var CmdRoute = []*api.Api{
	api.NewApi(&api.NewApiParamsType{
		Description: "执行命令",
		Path:        "/exec",
		Method:      gorequest.POST,
		Params:      controller.ExecParams{},
		Strategies: []api.StrategyData{
			{
				Strategy: global.BasicAuthStrategy,
			},
		},
		ControllerFunc: controller.CmdController.Exec,
	}),
}
