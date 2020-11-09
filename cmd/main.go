package main

import (
	"fmt"
	goWorkerPool "go-workerpool"
	"time"
)

type MyJob struct {
}

var count int

func (m MyJob) Do() error {
	count++
	fmt.Printf("do job [%d]\n", count)
	time.Sleep(50 * time.Millisecond)
	return nil
}

func main() {

	go func() {
		for {
			var job MyJob
			goWorkerPool.JobChannel <- job
			time.Sleep(10 * time.Millisecond)
		}
	}()

	dispatcher := goWorkerPool.NewDispatcher(10)
	dispatcher.Run()
}
