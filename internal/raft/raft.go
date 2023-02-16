package raft

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func New() Raft {
	return Raft{}
}

type Raft struct{}

func (r Raft) Start(ctx context.Context) {
	r.listNodes(ctx)
}

func (r Raft) listNodes(ctx context.Context) {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	for {
		// get pods in all the namespaces by omitting namespace
		// Or specify namespace to get pods in particular namespace
		pods, err := clientset.CoreV1().Pods("goraft").List(ctx, metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
		podsJson, err := json.MarshalIndent(pods.Items, "", "\t")
		if err != nil {
			continue
		}
		fmt.Printf("Pods are %s\n", string(podsJson))
		time.Sleep(10 * time.Second)
	}
}
