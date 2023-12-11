package workflow

import (
	"context"
	"fmt"
	workflow2 "github.com/argoproj/argo-workflows/v3/pkg/apiclient/workflow"

	"github.com/kubeTasker/kubeTaskerServer/rpc/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

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

func (l *ListWorkflowsLogic) ListWorkflows(in *core.WorkflowListRequest) (*core.WorkflowListRespond, error) {
	// todo: add your logic here and delete this line
	workflowListRequest := &workflow2.WorkflowListRequest{
		Namespace:   in.Namespace,
		ListOptions: in.ListOptions,
		Fields:      in.Fields,
	}

	resp, err := l.svcCtx.WorkflowClient.ListWorkflows(l.ctx, workflowListRequest)
	if err != nil {
		fmt.Println("获取工作流列表失败:", err)
	}

	// 处理工作流列表
	for _, wf := range resp.Items {
		fmt.Printf("工作流名称：%s，状态：%s\n", wf.ObjectMeta.Name, wf.Status.Phase)
	}
	fmt.Println(resp)
	if err != nil {
		return nil, err
	}

	return &core.WorkflowListRespond{
		WorkflowList: resp,
	}, nil
}
