package main

import (
	"github.com/vigneshv59/temporal-test/workflow"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
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

	w := worker.New(c, workflow.TaskQueue, worker.Options{
		MaxConcurrentActivityExecutionSize: 5,
		TaskQueueActivitiesPerSecond:       1,
	})

	w.RegisterWorkflow(workflow.TestWF)
	w.RegisterActivity(workflow.TestActivity)

	if err := w.Start(); err != nil {
		panic(err)
	}

	defer w.Stop()

	for range worker.InterruptCh() {
		return
	}
}
