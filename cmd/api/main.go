package main

import (
	"context"
	"fmt"

	"github.com/eduardoths/go-raft/internal/raft"
)

func main() {
	ctx := context.Background()
	raft, err := raft.New()
	if err != nil {
		panic(fmt.Errorf("Failed to instantiate raft: %v", err))
	}
	if err := raft.Start(ctx); err != nil {
		panic(fmt.Errorf("Failed to start raft: %v", err))
	}

	fmt.Println("Finished")
}
