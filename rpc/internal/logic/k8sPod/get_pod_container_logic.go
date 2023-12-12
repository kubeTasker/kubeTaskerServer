package k8sPod

import (
	"context"

	"github.com/kubeTasker/kubeTaskerServer/rpc/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPodContainerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPodContainerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPodContainerLogic {
	return &GetPodContainerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPodContainerLogic) GetPodContainer(in *core.GetPodDetailReq) (*core.GetPodDetailResp, error) {
	// todo: add your logic here and delete this line

	return &core.GetPodDetailResp{}, nil
}
