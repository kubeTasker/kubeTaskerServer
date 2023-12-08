package workflow

import (
	"context"

	"github.com/kubeTasker/kubeTaskerServer/rpc/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type WatchWorkflowsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewWatchWorkflowsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WatchWorkflowsLogic {
	return &WatchWorkflowsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *WatchWorkflowsLogic) WatchWorkflows(in *core.WatchWorkflowsRequest, stream core.Core_WatchWorkflowsServer) error {
	// todo: add your logic here and delete this line

	// Stream must supplement other logic by itself, and the logic is incomplete.
	// /root/TeamProject/kubeTaskerServer/rpc/types/core/core_grpc.pb.go

	return nil
}
