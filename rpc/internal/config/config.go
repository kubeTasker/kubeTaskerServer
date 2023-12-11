package config

import (
	"github.com/suyuan32/simple-admin-common/plugins/casbin"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"

	"github.com/suyuan32/simple-admin-common/config"
)

type Config struct {
	zrpc.RpcServerConf
	DatabaseConf    config.DatabaseConf
	CasbinConf      casbin.CasbinConf
	RedisConf       redis.RedisConf
	WorkflowConConf WorkflowConConf
	//WorkflowClientConf workflow.ClientConfig
}

type WorkflowConConf struct {
	Host string `json:",env=WORKFLOW_CON_CONF_HOST"` // *.*.*.*:*
	//Port int    `json:",env=WORKFLOW_CON_CONF_PORT"`
}
