package k8sdeployment

import (
	"context"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/kubeTasker/kubeTaskerServer/api/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDeployNumPerNpLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDeployNumPerNpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDeployNumPerNpLogic {
	return &GetDeployNumPerNpLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetDeployNumPerNpLogic) GetDeployNumPerNp(req *types.GetDeployNumPerNpReq) (resp *types.GetDeployNumPerNpResp, err error) {
	// todo: add your logic here and delete this line
	result, err := l.svcCtx.CoreRpc.GetDeployNumPerNp(l.ctx, &core.GetDeployNumPerNpReq{})
	if err != nil {
		return &types.GetDeployNumPerNpResp{
			Msg:  result.Msg,
			Data: nil,
		}, err
	}
	getDeployNumPerNpData := make([]*types.GetDeployNumPerNpData, 0)
	for _, v := range result.Data {
		getDeployNumPerNpData = append(getDeployNumPerNpData, &types.GetDeployNumPerNpData{
			Namespace:     v.Namespace,
			DeploymentNum: v.DeploymentNum,
		})
	}
	return &types.GetDeployNumPerNpResp{
		Msg:  result.Msg,
		Data: getDeployNumPerNpData,
	}, err
}
