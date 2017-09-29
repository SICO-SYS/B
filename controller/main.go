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
	"github.com/SiCo-Ops/cfg"
	"github.com/SiCo-Ops/dao/redis"
)

var (
	config     cfg.ConfigItems
	configPool = redis.NewPool()
	RPCServer  = grpc.NewServer()
)

const (
	configPath string = "config.json"
)

func ServePort() string {
	return config.RpcBPort
}

func init() {
	data, err := cfg.ReadFilePath(configPath)
	if err != nil {
		log.Fatalln(err)
	}
	cfg.Unmarshal(data, &config)

	configPool = redis.InitPool(config.RedisConfigHost, config.RedisConfigPort, config.RedisConfigAuth)
	rediserr := redis.Hmset(configPool, "system.config", &config)
	if rediserr != nil {
		raven.CaptureError(rediserr, nil)
		log.Fatalln(rediserr)
	}

	pb.RegisterConfigServiceServer(RPCServer, &ConfigService{})

	if config.SentryBStatus == "active" && config.SentryBDSN != "" {
		raven.SetDSN(config.SentryBDSN)
	}
}
