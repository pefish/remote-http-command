package global

import api_strategy "github.com/pefish/go-core-strategy/api-strategy"

type Config struct {
	ServerHost string `json:"server-host"`
	ServerPort uint64 `json:"server-port"`
	Username   string `json:"username"`
	Password   string `json:"password"`
}

var GlobalConfig Config
var BasicAuthStrategy *api_strategy.BasicAuthStrategy = api_strategy.NewBasicAuthStrategy()
