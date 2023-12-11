package workflow

import (
	"context"
	workflow2 "github.com/argoproj/argo-workflows/v3/pkg/apiclient/workflow"

	"github.com/kubeTasker/kubeTaskerServer/rpc/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateWorkflowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateWorkflowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateWorkflowLogic {
	return &CreateWorkflowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Workflow management
func (l *CreateWorkflowLogic) CreateWorkflow(in *core.WorkflowCreateRequest) (*core.WorkflowRespond, error) {
	// todo: add your logic here and delete this line
	workflowCreateRequest := &workflow2.WorkflowCreateRequest{
		Namespace:     in.Namespace,
		Workflow:      in.Workflow,
		InstanceID:    in.InstanceID,
		ServerDryRun:  in.ServerDryRun,
		CreateOptions: in.CreateOptions,
	}

	resp, err := l.svcCtx.WorkflowClient.CreateWorkflow(l.ctx, workflowCreateRequest)

	if err != nil {
		return nil, err
	}

	return &core.WorkflowRespond{
		Workflow: resp,
	}, nil
}
