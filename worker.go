package go_workerpool

import "log"

type Job interface {
	Do() error
}

//job 队列
var JobChannel = make(chan Job)

type Worker struct {
	WorkerPool chan chan Job
	JobChannel chan Job
	done       chan bool
}

func NewWorker(pool chan chan Job) *Worker {
	return &Worker{
		WorkerPool: pool,
		JobChannel: make(chan Job),
		done:       make(chan bool),
	}
}

func (w *Worker) Start() {
	go func() {
		for {
			//将自己的JobChannel注册到 WorkerPool中去
			w.WorkerPool <- w.JobChannel
			select {
			case job := <-w.JobChannel:
				err := job.Do()
				if err != nil {
					log.Printf("job do error:%v\n", err)
				}
			case <-w.done:
				return
			}
		}
	}()
}

func (w *Worker) Stop() {
	go func() {
		w.done <- true
	}()
}
