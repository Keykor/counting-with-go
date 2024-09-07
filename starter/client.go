package starter

import (
	"context"
	"go.temporal.io/sdk/client"
	"log"
	"counting/counter"
)

func UpdateCounter(amount int) {
	cli, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}

	options := client.StartWorkflowOptions{
		TaskQueue: "counter-worker",
	}

	go func() {
		we, err := cli.ExecuteWorkflow(context.Background(), options, counter.UpdateCounterWorkFlow, amount)
		if err != nil {
			log.Fatalln("Unable to execute workflow", err)
		}

		err = we.Get(context.Background(), nil)
		if err != nil {
			log.Fatalln("Execution failed", err)
		}

	    defer cli.Close()
	}()
}

func GetCounter() int {
    cli, err := client.Dial(client.Options{})
    defer cli.Close()

    if err != nil {
        log.Fatalln("Unable to create client", err)
    }

    options := client.StartWorkflowOptions{
        TaskQueue: "counter-worker",
    }

    var result int

    we, err := cli.ExecuteWorkflow(context.Background(), options, counter.GetCounterWorkFlow)
    if err != nil {
        log.Fatalln("Unable to execute workflow", err)
    }
    
    err = we.Get(context.Background(), &result)
    if err != nil {
        log.Fatalln("Execution failed", err)
    }

    return result
}

func ResetCounter() {
    cli, err := client.Dial(client.Options{})
    if err != nil {
        log.Fatalln("Unable to create client", err)
    }

    options := client.StartWorkflowOptions{
        TaskQueue: "counter-worker",
    }

    go func() {
        we, err := cli.ExecuteWorkflow(context.Background(), options, counter.ResetCounterWorkFlow)
        if err != nil {
            log.Fatalln("Unable to execute workflow", err)
        }
    
        err = we.Get(context.Background(), nil)
        if err != nil {
            log.Fatalln("Execution failed", err)
        }

        defer cli.Close()
    }()
}
