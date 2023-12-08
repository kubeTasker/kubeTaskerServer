package workflow

import (
	"context"
	workflow2 "github.com/argoproj/argo-workflows/v3/pkg/apiclient/workflow"

	"github.com/kubeTasker/kubeTaskerServer/rpc/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type RetryWorkflowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRetryWorkflowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RetryWorkflowLogic {
	return &RetryWorkflowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RetryWorkflowLogic) RetryWorkflow(in *core.WorkflowRetryRequest) (*core.WorkflowRespond, error) {
	// todo: add your logic here and delete this line
	workflowRetryRequest := &workflow2.WorkflowRetryRequest{
		Name:              in.Name,
		Namespace:         in.Namespace,
		RestartSuccessful: in.RestartSuccessful,
		NodeFieldSelector: in.NodeFieldSelector,
		Parameters:        in.Parameters,
	}

	resp, err := l.svcCtx.WorkflowClient.RetryWorkflow(l.ctx, workflowRetryRequest)

	if err != nil {
		return nil, err
	}

	return &core.WorkflowRespond{
		Workflow: resp,
	}, nil
}
