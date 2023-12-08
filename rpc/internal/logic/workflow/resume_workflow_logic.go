package workflow

import (
	"context"
	workflow2 "github.com/argoproj/argo-workflows/v3/pkg/apiclient/workflow"

	"github.com/kubeTasker/kubeTaskerServer/rpc/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type ResumeWorkflowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewResumeWorkflowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResumeWorkflowLogic {
	return &ResumeWorkflowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ResumeWorkflowLogic) ResumeWorkflow(in *core.WorkflowResumeRequest) (*core.WorkflowRespond, error) {
	// todo: add your logic here and delete this line
	workflowResumeRequest := &workflow2.WorkflowResumeRequest{
		Name:              in.Name,
		Namespace:         in.Namespace,
		NodeFieldSelector: in.NodeFieldSelector,
	}

	resp, err := l.svcCtx.WorkflowClient.ResumeWorkflow(l.ctx, workflowResumeRequest)

	if err != nil {
		return nil, err
	}

	return &core.WorkflowRespond{
		Workflow: resp,
	}, nil
}
