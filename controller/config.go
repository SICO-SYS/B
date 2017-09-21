/*

LICENSE:  MIT
Author:   sine
Email:    sinerwr@gmail.com

Contributors

*/

package controller

import (
	"encoding/json"
	"golang.org/x/net/context"

	"github.com/SiCo-Ops/Pb"
	"github.com/SiCo-Ops/dao/redis"
)

type ConfigService struct{}

func (c *ConfigService) PushRPC(ctx context.Context, in *pb.ConfigPushCall) (*pb.ConfigPushBack, error) {
	v := make(map[string]interface{})
	json.Unmarshal(in.Params, &v)
	if in.Environment == "" {
		in.Environment = "default"
	}
	key := in.Id + "." + in.Environment
	err := redis.Hmset(configPool, key, v)
	if err != nil {
		return &pb.ConfigPushBack{Code: 1}, err
	}
	return &pb.ConfigPushBack{Code: 0}, nil
}

func (c *ConfigService) PullRPC(ctx context.Context, in *pb.ConfigPullCall) (*pb.ConfigPullBack, error) {
	if in.Environment == "" {
		in.Environment = "default"
	}
	key := in.Id + "." + in.Environment
	result, err := redis.Hgetall(configPool, key)
	if err != nil {
		return &pb.ConfigPullBack{Code: 1}, err
	}
	data, err := json.Marshal(result)
	if err != nil {
		return &pb.ConfigPullBack{Code: 1}, err
	}
	return &pb.ConfigPullBack{Code: 0, Params: data}, nil
}
