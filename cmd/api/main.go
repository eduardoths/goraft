package main

import (
	"context"
	"fmt"

	"github.com/eduardoths/go-raft/internal/raft"
)

func main() {
	ctx := context.Background()
	raft := raft.New()
	fmt.Println("Starting raft")
	raft.Start(ctx)
	fmt.Println("Finished raft")
}
