package raft

import (
	"context"
	"encoding/json"
	"fmt"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func New() (Raft, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return Raft{}, err
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return Raft{}, err
	}
	return Raft{
		clientset: clientset,
	}, nil
}

type Raft struct {
	clientset *kubernetes.Clientset
}

func (r Raft) Start(ctx context.Context) error {
	pods, err := r.listPods(ctx)
	if err != nil {
		return err
	}
	podsJson, err := json.MarshalIndent(pods, "", "  ")
	if err != nil {
		return err
	}

	fmt.Printf("Found %d pods:\n%v\n", len(pods), podsJson)
	return nil
}

func (r Raft) listPods(ctx context.Context) ([]v1.Pod, error) {
	pods, err := r.clientset.CoreV1().Pods("goraft").List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return pods.Items, nil
}
