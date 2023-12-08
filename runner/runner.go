package main

import (
	"context"
	"github.com/google/uuid"
	"github.com/vigneshv59/temporal-test/workflow"
	"go.temporal.io/sdk/client"
	"os"
)

func main() {
	opts := client.Options{
		HostPort: os.Getenv("TEMPORAL"),
	}
	c, err := client.Dial(opts)

	if err != nil {
		panic(err)
	}

	if _, err := c.ExecuteWorkflow(context.Background(), client.StartWorkflowOptions{
		ID:        uuid.NewString(),
		TaskQueue: workflow.TaskQueue,
	}, workflow.TestWF); err != nil {
		panic(err)
	}
}
