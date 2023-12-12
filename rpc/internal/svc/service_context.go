package svc

import (
	"github.com/kubeTasker/kubeTaskerServer/rpc/ent"
	"github.com/kubeTasker/kubeTaskerServer/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"

	"context"
	_ "github.com/kubeTasker/kubeTaskerServer/rpc/ent/runtime"
)

type ServiceContext struct {
	context.Context
	Config config.Config
	DB     *ent.Client
	Redis  *redis.Redis
}

func NewServiceContext(c config.Config) (ctx *ServiceContext, err error) {
	db := ent.NewClient(
		ent.Log(logx.Info), // logger
		ent.Driver(c.DatabaseConf.NewNoCacheDriver()),
		ent.Debug(), // debug mode
	)

	if err != nil {
		return nil, err
	}

	return &ServiceContext{
		Config: c,
		DB:     db,
		Redis:  redis.MustNewRedis(c.RedisConf),
	}, nil

}
