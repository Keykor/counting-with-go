package worker

import (
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"log"
	"counting/counter"
)

func CreateWorker() {
	client, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}

	w := worker.New(client, "counter-worker", worker.Options{})
	w.RegisterWorkflow(counter.UpdateCounterWorkFlow)
    w.RegisterActivity(counter.UpdateCounterActivity)
    w.RegisterWorkflow(counter.GetCounterWorkFlow)
    w.RegisterActivity(counter.GetCounterActivity)
    w.RegisterWorkflow(counter.ResetCounterWorkFlow)
	w.RegisterActivity(counter.ResetCounterActivity)

	go func() {
		err = w.Run(nil)
		if err != nil {
			log.Fatalln("Unable to start worker", err)
		}
	    defer client.Close()
	}()
}
