package config

import (
	"p_meego/server/shared/kitex_gen/user/userservice"
)

var (
	IsDev bool

	// 持有通过consul获取的全部服务配置，包括本服务的ip、端口、providerName
	GlobalServerConfig ServerConfig

	GlobalLocalConfig LocalConfig

	GlobalUserProviderClient userservice.Client
)
