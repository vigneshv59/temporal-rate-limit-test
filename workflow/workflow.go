package workflow

import (
	"fmt"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
	"time"
)

const TaskQueue string = "rl_queue"

func TestWF(ctx workflow.Context) error {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute * 2,
		RetryPolicy: &temporal.RetryPolicy{
			MaximumAttempts: 1,
		},
		WaitForCancellation: true,
		TaskQueue:           TaskQueue,
	}

	ctx = workflow.WithActivityOptions(ctx, options)
	futures := make([]workflow.Future, 0)

	for i := 0; i < 1000; i++ {
		f := workflow.ExecuteActivity(
			ctx,
			TestActivity,
			nil,
		)
		futures = append(futures, f)
	}

	for _, f := range futures {
		f.Get(ctx, nil)
	}

	return nil
}

func TestActivity() error {
	fmt.Println("Test activity", time.Now())
	return nil
}
