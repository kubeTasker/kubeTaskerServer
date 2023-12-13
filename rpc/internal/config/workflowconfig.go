package config

import (
	"log"
	"os"

	wfclientset "github.com/kubeTasker/kubeTasker/pkg/client/clientset/versioned"
	"github.com/kubeTasker/kubeTasker/pkg/client/clientset/versioned/typed/workflow/v1alpha1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type WorkflowConfig struct {
	RestConfig   *rest.Config
	ClientConfig clientcmd.ClientConfig
	WFClient     v1alpha1.WorkflowInterface
	Clientset    *kubernetes.Clientset
}

func InitKubeClientConfig(wfconfig *WorkflowConfig) {
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()

	overrides := clientcmd.ConfigOverrides{}

	wfconfig.ClientConfig = clientcmd.NewInteractiveDeferredLoadingClientConfig(loadingRules, &overrides, os.Stdin)
}

func InitKubeClient(wfconfig *WorkflowConfig) *kubernetes.Clientset {
	if wfconfig.Clientset != nil {
		return wfconfig.Clientset
	}
	var err error
	wfconfig.RestConfig, err = wfconfig.ClientConfig.ClientConfig()
	if err != nil {
		log.Fatal(err)
	}

	wfconfig.Clientset, err = kubernetes.NewForConfig(wfconfig.RestConfig)
	if err != nil {
		log.Fatal(err)
	}
	return wfconfig.Clientset
}

func InitWorkflowClient(wfconfig *WorkflowConfig, ns ...string) {
	if wfconfig.WFClient != nil {
		return
	}
	InitKubeClient(wfconfig)
	var namespace string
	var err error
	if len(ns) > 0 {
		namespace = ns[0]
	} else {
		namespace, _, err = wfconfig.ClientConfig.Namespace()
		if err != nil {
			log.Fatal(err)
		}
	}
	wfcs := wfclientset.NewForConfigOrDie(wfconfig.RestConfig)
	wfconfig.WFClient = wfcs.KubetaskerV1alpha1().Workflows(namespace)
}

func NewWorkflowConfig() *WorkflowConfig {
	var wfconfig WorkflowConfig
	InitKubeClientConfig(&wfconfig)
	InitWorkflowClient(&wfconfig)
	return &wfconfig
}
