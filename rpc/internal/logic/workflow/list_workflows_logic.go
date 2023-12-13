package workflow

import (
	"context"

	"github.com/kubeTasker/kubeTaskerServer/rpc/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"
	"k8s.io/apimachinery/pkg/labels"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListWorkflowsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListWorkflowsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListWorkflowsLogic {
	return &ListWorkflowsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListWorkflowsLogic) ListWorkflows(in *core.WorkflowListRequest) (*core.WorkflowListResponse, error) {
	labelSelector := labels.NewSelector()
	listOpts := *in.ListOptions
	listOpts.LabelSelector = labelSelector.String()
	
	wfList, err := l.svcCtx.WFConfig.WFClient.List(context.TODO(), listOpts)
	if err != nil {
		return &core.WorkflowListResponse{}, nil
	}
	return &core.WorkflowListResponse{Workflowlist: wfList}, nil
}
