package main

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"p_meego/server/cmd/api/config"
	"p_meego/server/cmd/user/initialize"
	"p_meego/server/cmd/user/pkg/mysql"
	"p_meego/server/shared/kitex_gen/user/userservice"
)

func main() {
	// 本地配置文件初始化
	initialize.InitLocalConfig()
	// 根据本地配置文件初始化日志
	initialize.InitLogger()
	// 根据本地配置文件初始化配置
	initialize.InitConfig()
	// 根据本地配置文件初始化数据库
	db := initialize.InitDB()
	// 注册中心
	r, info := initialize.InitRegistry(config.GlobalServerConfig.Port)
	//p := provider.NewOpenTelemetryProvider(
	//	provider.WithServiceName(config.GlobalServerConfig.Name),
	//	//provider.WithExportEndpoint(config.GlobalServerConfig.OtelInfo.EndPoint),
	//	provider.WithInsecure(),
	//)
	//defer p.Shutdown(context.Background())

	svr := userservice.NewServer(&UserServiceImpl{UserContext: mysql.Use(db)},
		server.WithRegistry(r),
		server.WithRegistryInfo(info),
		server.WithLimit(&limit.Option{MaxConnections: 2000, MaxQPS: 500}),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: config.GlobalServerConfig.Name}))

	err := svr.Run()

	if err != nil {
		klog.Fatal(err.Error())
	}
}
