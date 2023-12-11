package workflow

import (
	"context"
	workflow2 "github.com/argoproj/argo-workflows/v3/pkg/apiclient/workflow"

	"github.com/kubeTasker/kubeTaskerServer/rpc/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type StopWorkflowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewStopWorkflowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StopWorkflowLogic {
	return &StopWorkflowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *StopWorkflowLogic) StopWorkflow(in *core.WorkflowStopRequest) (*core.WorkflowRespond, error) {
	// todo: add your logic here and delete this line
	workflowStopRequest := &workflow2.WorkflowStopRequest{
		Name:              in.Name,
		Namespace:         in.Namespace,
		NodeFieldSelector: in.NodeFieldSelector,
		Message:           in.Message,
	}

	resp, err := l.svcCtx.WorkflowClient.StopWorkflow(l.ctx, workflowStopRequest)

	if err != nil {
		return nil, err
	}

	return &core.WorkflowRespond{
		Workflow: resp,
	}, nil
}
