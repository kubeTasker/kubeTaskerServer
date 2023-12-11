package workflow

import (
	"context"
	workflow2 "github.com/argoproj/argo-workflows/v3/pkg/apiclient/workflow"

	"github.com/kubeTasker/kubeTaskerServer/rpc/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteWorkflowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteWorkflowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteWorkflowLogic {
	return &DeleteWorkflowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteWorkflowLogic) DeleteWorkflow(in *core.WorkflowDeleteRequest) (*core.WorkflowDeleteResponse, error) {
	// todo: add your logic here and delete this line
	workflowDeleteRequest := &workflow2.WorkflowDeleteRequest{
		Name:          in.Name,
		Namespace:     in.Namespace,
		DeleteOptions: in.DeleteOptions,
		Force:         in.Force,
	}

	_, err := l.svcCtx.WorkflowClient.DeleteWorkflow(l.ctx, workflowDeleteRequest)

	if err != nil {
		return nil, err
	}

	return &core.WorkflowDeleteResponse{}, nil
}
