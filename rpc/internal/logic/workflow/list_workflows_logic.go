package workflow

import (
	"context"
	workflow2 "github.com/argoproj/argo-workflows/v3/pkg/apiclient/workflow"

	"github.com/kubeTasker/kubeTaskerServer/rpc/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListWorkflowsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListWorkflowsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListWorkflowsLogic {
	return &ListWorkflowsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListWorkflowsLogic) ListWorkflows(in *core.WorkflowListRequest) (*core.WorkflowListRespond, error) {
	// todo: add your logic here and delete this line
	workflowListRequest := &workflow2.WorkflowListRequest{
		Namespace:   in.Namespace,
		ListOptions: in.ListOptions,
		Fields:      in.Fields,
	}

	resp, err := l.svcCtx.WorkflowClient.ListWorkflows(l.ctx, workflowListRequest)

	if err != nil {
		return nil, err
	}

	return &core.WorkflowListRespond{
		WorkflowList: resp,
	}, nil
}
