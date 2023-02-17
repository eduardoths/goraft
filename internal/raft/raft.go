package raft

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/eduardoths/go-raft/internal/utils/kubernetes"
	"github.com/eduardoths/go-raft/internal/utils/random"
)

const (
	MIN_HEARTBEAT_MS = 150
	MAX_HEARTBEAT_MS = 300
)

func New() (Raft, error) {
	return Raft{
		heartbeat: random.Heartbeat(MIN_HEARTBEAT_MS, MAX_HEARTBEAT_MS),
		kubeUtils: kubernetes.NewK8sUtils(),
	}, nil
}

type Raft struct {
	heartbeat time.Duration
	kubeUtils kubernetes.KubeUtils
}

func (r Raft) Start(ctx context.Context) error {
	pods, err := r.kubeUtils.ListPods(ctx)
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
