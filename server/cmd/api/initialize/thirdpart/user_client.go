package thirdpart

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/loadbalance"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	consul "github.com/kitex-contrib/registry-consul"
	"p_meego/server/cmd/api/config"
	"p_meego/server/shared/kitex_gen/user/userservice"
)

func initUser() {
	// init resolver
	r, err := consul.NewConsulResolver(fmt.Sprintf("%s:%d",
		config.GlobalLocalConfig.ConsulConfig.Host,
		config.GlobalLocalConfig.ConsulConfig.Port))
	if err != nil {
		hlog.Fatalf("new consul client failed: %s", err.Error())
	}

	// create a new client
	c, err := userservice.NewClient(
		config.GlobalServerConfig.UserProviderService.Name,
		client.WithResolver(r),                                     // service discovery
		client.WithLoadBalancer(loadbalance.NewWeightedBalancer()), // load balance
		client.WithMuxConnection(1),                                // multiplexing
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: config.GlobalServerConfig.UserProviderService.Name}),
	)
	if err != nil {
		hlog.Fatalf("ERROR: cannot init client: %v\n", err)
	}
	config.GlobalUserProviderClient = c
}
