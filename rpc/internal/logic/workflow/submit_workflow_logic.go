package workflow

import (
	"context"
	workflow2 "github.com/argoproj/argo-workflows/v3/pkg/apiclient/workflow"

	"github.com/kubeTasker/kubeTaskerServer/rpc/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubmitWorkflowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSubmitWorkflowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubmitWorkflowLogic {
	return &SubmitWorkflowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SubmitWorkflowLogic) SubmitWorkflow(in *core.WorkflowSubmitRequest) (*core.WorkflowRespond, error) {
	// todo: add your logic here and delete this line
	workflowSubmitRequest := &workflow2.WorkflowSubmitRequest{
		Namespace:     in.Namespace,
		ResourceKind:  in.ResourceKind,
		ResourceName:  in.ResourceName,
		SubmitOptions: in.SubmitOptions,
	}

	resp, err := l.svcCtx.WorkflowClient.SubmitWorkflow(l.ctx, workflowSubmitRequest)

	if err != nil {
		return nil, err
	}

	return &core.WorkflowRespond{
		Workflow: resp,
	}, nil
}
