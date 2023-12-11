package workflow

import (
	"context"
	workflow2 "github.com/argoproj/argo-workflows/v3/pkg/apiclient/workflow"

	"github.com/kubeTasker/kubeTaskerServer/rpc/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type SuspendWorkflowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSuspendWorkflowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SuspendWorkflowLogic {
	return &SuspendWorkflowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SuspendWorkflowLogic) SuspendWorkflow(in *core.WorkflowSuspendRequest) (*core.WorkflowRespond, error) {
	// todo: add your logic here and delete this line
	workflowSuspendRequest := &workflow2.WorkflowSuspendRequest{
		Name:      in.Name,
		Namespace: in.Namespace,
	}

	resp, err := l.svcCtx.WorkflowClient.SuspendWorkflow(l.ctx, workflowSuspendRequest)

	if err != nil {
		return nil, err
	}

	return &core.WorkflowRespond{
		Workflow: resp,
	}, nil
}
