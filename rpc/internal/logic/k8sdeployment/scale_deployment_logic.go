package k8sdeployment

import (
	"context"
	"github.com/kubeTasker/kubeTaskerServer/rpc/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type ScaleDeploymentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewScaleDeploymentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ScaleDeploymentLogic {
	return &ScaleDeploymentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 设置deployment副本数
func (l *ScaleDeploymentLogic) ScaleDeployment(in *core.ScaleDeploymentReq) (*core.ScaleDeploymentResp, error) {
	// todo: add your logic here and delete this line
	deployment := &Deployment{}
	resp, _ := deployment.ScaleDeployment(l, in)
	return resp, nil
}
