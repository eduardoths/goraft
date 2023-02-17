package kubernetes

import (
	"context"
	"log"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type KubeUtils struct {
	clientset *kubernetes.Clientset
}

func NewK8sUtils() KubeUtils {
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Panicf("failed to start k8s config: %v", err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Panicf("failed to start k8s clientset: %v", err)
	}
	return KubeUtils{
		clientset: clientset,
	}
}

func (ku KubeUtils) ListPods(ctx context.Context) ([]v1.Pod, error) {
	pods, err := ku.clientset.CoreV1().Pods("goraft").List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return pods.Items, nil
}
