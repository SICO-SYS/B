/*

LICENSE:  MIT
Author:   sine
Email:    sinerwr@gmail.com

*/

package controller

import (
	"github.com/getsentry/raven-go"
	"google.golang.org/grpc"
	"log"

	"github.com/SiCo-Ops/Pb"
	"github.com/SiCo-Ops/cfg/v2"
	"github.com/SiCo-Ops/dao/redis"
)

var (
	config     cfg.ConfigItems
	configPool = redis.NewPool()
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

	configPool = redis.InitPool(config.RedisConfigHost, config.RedisConfigPort, config.RedisConfigAuth)
	err := redis.Hmset(configPool, "system.config", &config)
	if err != nil {
		raven.CaptureError(err, nil)
		log.Fatalln(err)
	}

	pb.RegisterConfigServiceServer(RPCServer, &ConfigService{})

	if config.SentryBStatus == "active" && config.SentryBDSN != "" {
		raven.SetDSN(config.SentryBDSN)
	}
}
