package go_workerpool

type Dispatcher struct {
	WorkerPool chan chan Job
	Len        int
}

func NewDispatcher(n int) *Dispatcher {
	return &Dispatcher{
		WorkerPool: make(chan chan Job, n),
		Len:        n,
	}
}

func (d *Dispatcher) Run() {
	for i := 0; i < d.Len; i++ {
		worker := NewWorker(d.WorkerPool)
		worker.Start()
	}

}

func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-JobChannel:
			go func(job Job) {
				jobChannel := <-d.WorkerPool
				jobChannel <- job
			}(job)
		}
	}
}
