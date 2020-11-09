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
	d.dispatch()
}

func (d *Dispatcher) dispatch() {
	//这地方改成队列
	var jobQ []Job
	var jobChannelQ []chan Job
	for {
		var activeJob Job
		var activeJobChannel chan Job
		if len(jobQ) > 0 && len(jobChannelQ) > 0 {
			activeJob = jobQ[0]
			activeJobChannel = jobChannelQ[0]
		}
		select {
		case job := <-JobChannel:
			jobQ = append(jobQ, job)
			//fmt.Printf("receive job, current jobQ size is %d\n", len(jobQ))
		case jobChannel := <-d.WorkerPool:
			jobChannelQ = append(jobChannelQ, jobChannel)
			//fmt.Printf("receive ready worker, current jobChannelQ size is %d\n", len(jobChannelQ))
		case activeJobChannel <- activeJob:
			jobQ = jobQ[1:]
			jobChannelQ = jobChannelQ[1:]
		}
	}
}
