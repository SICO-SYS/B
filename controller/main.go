/*

LICENSE:  MIT
Author:   sine
Email:    sinerwr@gmail.com

*/

package controller

import (
	// "github.com/getsentry/raven-go"
	"google.golang.org/grpc"

	"github.com/SiCo-Ops/Pb"
	"github.com/SiCo-Ops/cfg/v2"
	"github.com/SiCo-Ops/dao/redis"
)

var (
	config     cfg.ConfigItems
	configPool = redis.Pool("", "", "")
	RPCServer  = grpc.NewServer()
)

func ServePort() string {
	return config.RpcBPort
}

func init() {
	data := cfg.ReadLocalFile()

	if data != nil {
		cfg.Unmarshal(data, &config)
	}

	configPool = redis.Pool(config.RedisConfigHost, config.RedisConfigPort, config.RedisConfigAuth)
	redis.Hmset(configPool, "system.config", &config)

	pb.RegisterConfigServiceServer(RPCServer, &ConfigService{})
}
