package workflow

import (
	"context"
	workflow2 "github.com/argoproj/argo-workflows/v3/pkg/apiclient/workflow"

	"github.com/kubeTasker/kubeTaskerServer/rpc/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type LintWorkflowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLintWorkflowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LintWorkflowLogic {
	return &LintWorkflowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LintWorkflowLogic) LintWorkflow(in *core.WorkflowLintRequest) (*core.WorkflowRespond, error) {
	// todo: add your logic here and delete this line
	workflowLintRequest := &workflow2.WorkflowLintRequest{
		Namespace: in.Namespace,
		Workflow:  in.Workflow,
	}

	resp, err := l.svcCtx.WorkflowClient.LintWorkflow(l.ctx, workflowLintRequest)

	if err != nil {
		return nil, err
	}

	return &core.WorkflowRespond{
		Workflow: resp,
	}, nil
}
