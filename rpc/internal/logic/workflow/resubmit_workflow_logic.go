package workflow

import (
	"context"
	workflow2 "github.com/argoproj/argo-workflows/v3/pkg/apiclient/workflow"
	"github.com/kubeTasker/kubeTaskerServer/rpc/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type ResubmitWorkflowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewResubmitWorkflowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResubmitWorkflowLogic {
	return &ResubmitWorkflowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ResubmitWorkflowLogic) ResubmitWorkflow(in *core.WorkflowResubmitRequest) (*core.WorkflowRespond, error) {
	// todo: add your logic here and delete this line
	workflowResubmitRequest := &workflow2.WorkflowResubmitRequest{
		Name:       in.Name,
		Namespace:  in.Namespace,
		Memoized:   in.Memoized,
		Parameters: in.Parameters,
	}

	resp, err := l.svcCtx.WorkflowClient.ResubmitWorkflow(l.ctx, workflowResubmitRequest)

	if err != nil {
		return nil, err
	}

	return &core.WorkflowRespond{
		Workflow: resp,
	}, nil
}
