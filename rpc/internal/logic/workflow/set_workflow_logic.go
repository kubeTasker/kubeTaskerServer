package workflow

import (
	"context"
	workflow2 "github.com/argoproj/argo-workflows/v3/pkg/apiclient/workflow"

	"github.com/kubeTasker/kubeTaskerServer/rpc/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetWorkflowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetWorkflowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetWorkflowLogic {
	return &SetWorkflowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetWorkflowLogic) SetWorkflow(in *core.WorkflowSetRequest) (*core.WorkflowRespond, error) {
	// todo: add your logic here and delete this line
	workflowSetRequest := &workflow2.WorkflowSetRequest{
		Name:              in.Name,
		Namespace:         in.Namespace,
		NodeFieldSelector: in.NodeFieldSelector,
		Message:           in.Message,
		Phase:             in.Phase,
		OutputParameters:  in.OutputParameters,
	}

	resp, err := l.svcCtx.WorkflowClient.SetWorkflow(l.ctx, workflowSetRequest)

	if err != nil {
		return nil, err
	}

	return &core.WorkflowRespond{
		Workflow: resp,
	}, nil
}
