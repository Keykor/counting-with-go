package counter

import (
    "go.temporal.io/sdk/workflow"
    "time"
)

func UpdateCounterWorkFlow(ctx workflow.Context, amount int) error {
    activityOptions := workflow.ActivityOptions{
        StartToCloseTimeout: time.Minute,
    }
    ctx = workflow.WithActivityOptions(ctx, activityOptions)
    
    futureList := make([]workflow.Future, amount)

    for i := 0; i < amount; i++ {
        futureList[i] = workflow.ExecuteActivity(ctx, UpdateCounterActivity, i)
    }

    for i := 0; i < amount; i++ {
        err := futureList[i].Get(ctx, nil)
        if err != nil {
            return err
        }
    }
    
    return nil
}

func GetCounterWorkFlow(ctx workflow.Context) (int, error) {
    activityOptions := workflow.ActivityOptions{
        StartToCloseTimeout: time.Minute,
    }
    ctx = workflow.WithActivityOptions(ctx, activityOptions)
    
    var counter int
    err := workflow.ExecuteActivity(ctx, GetCounterActivity).Get(ctx, &counter)
    if err != nil {
        return 0, err
    }
    
    return counter, nil
}

func ResetCounterWorkFlow(ctx workflow.Context) error {
    activityOptions := workflow.ActivityOptions{
        StartToCloseTimeout: time.Minute,
    }
    ctx = workflow.WithActivityOptions(ctx, activityOptions)
    
    err := workflow.ExecuteActivity(ctx, ResetCounterActivity).Get(ctx, nil)
    return err
}
