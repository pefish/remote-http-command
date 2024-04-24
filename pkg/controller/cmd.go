package controller

import (
	_type "github.com/pefish/go-core-type/api-session"
	go_error "github.com/pefish/go-error"
	go_logger "github.com/pefish/go-logger"
	go_shell "github.com/pefish/go-shell"
)

type CmdControllerType struct {
}

var CmdController = CmdControllerType{}

type ExecParams struct {
	Cmd string `json:"cmd" validate:"required"`
}

func (c *CmdControllerType) Exec(apiSession _type.IApiSession) (interface{}, *go_error.ErrorInfo) {
	var params ExecParams
	err := apiSession.ScanParams(&params)
	if err != nil {
		go_logger.Logger.ErrorF("Read params error. %+v", err)
		return nil, go_error.INTERNAL_ERROR
	}

	execResult, err := go_shell.ExecForResult(go_shell.NewCmd(params.Cmd))
	if err != nil {
		return nil, go_error.WrapWithErr(err)
	}
	return execResult, nil
}
