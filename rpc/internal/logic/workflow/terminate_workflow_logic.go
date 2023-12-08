package workflow

import (
	"context"
	workflow2 "github.com/argoproj/argo-workflows/v3/pkg/apiclient/workflow"

	"github.com/kubeTasker/kubeTaskerServer/rpc/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type TerminateWorkflowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTerminateWorkflowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TerminateWorkflowLogic {
	return &TerminateWorkflowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TerminateWorkflowLogic) TerminateWorkflow(in *core.WorkflowTerminateRequest) (*core.WorkflowRespond, error) {
	// todo: add your logic here and delete this line
	workflowTerminateRequest := &workflow2.WorkflowTerminateRequest{
		Name:      in.Name,
		Namespace: in.Namespace,
	}

	resp, err := l.svcCtx.WorkflowClient.TerminateWorkflow(l.ctx, workflowTerminateRequest)

	if err != nil {
		return nil, err
	}

	return &core.WorkflowRespond{
		Workflow: resp,
	}, nil
}
