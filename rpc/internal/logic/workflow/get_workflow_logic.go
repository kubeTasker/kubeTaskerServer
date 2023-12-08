package workflow

import (
	"context"
	workflow2 "github.com/argoproj/argo-workflows/v3/pkg/apiclient/workflow"

	"github.com/kubeTasker/kubeTaskerServer/rpc/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetWorkflowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetWorkflowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetWorkflowLogic {
	return &GetWorkflowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetWorkflowLogic) GetWorkflow(in *core.WorkflowGetRequest) (*core.WorkflowRespond, error) {
	// todo: add your logic here and delete this line
	workflowGetRequest := &workflow2.WorkflowGetRequest{
		Name:       in.Name,
		Namespace:  in.Namespace,
		GetOptions: in.GetOptions,
		Fields:     in.Fields,
	}

	resp, err := l.svcCtx.WorkflowClient.GetWorkflow(l.ctx, workflowGetRequest)

	if err != nil {
		return nil, err
	}

	return &core.WorkflowRespond{
		Workflow: resp,
	}, nil
}
