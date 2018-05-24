package test

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
)

var (
	//Max_Num = os.Getenv("MAX_NUM")
	MaxWorker = runtime.NumCPU()
	MaxQueue  = 1000
)

type Serload struct {
	pri string
}

type Job struct {
	serload Serload
}

var JobQueue chan Job

type Worker struct {
	WorkerPool chan chan Job
	JobChannel chan Job
	Quit       chan bool
}

func NewWorker(workPool chan chan Job) Worker {
	return Worker{
		WorkerPool: workPool,
		JobChannel: make(chan Job),
		Quit:       make(chan bool),
	}
}

func (w Worker) Start() {
	go func() {
		for {
			w.WorkerPool <- w.JobChannel
			select {
			case job := <-w.JobChannel:
				// excute job
				fmt.Println(job.serload.pri)
			case <-w.Quit:
				return
			}
		}
	}()
}

func (w Worker) Stop() {
	go func() {
		w.Quit <- true
	}()
}

type Dispatcher struct {
	MaxWorkers int
	WorkerPool chan chan Job
	Quit       chan bool
}

func NewDispatcher(maxWorkers int) *Dispatcher {
	pool := make(chan chan Job, maxWorkers)
	return &Dispatcher{MaxWorkers: maxWorkers, WorkerPool: pool, Quit: make(chan bool)}
}

func (d *Dispatcher) Run() {
	for i := 0; i < d.MaxWorkers; i++ {
		worker := NewWorker(d.WorkerPool)
		worker.Start()
	}

	go d.Dispatch()
}

func (d *Dispatcher) Stop() {
	go func() {
		d.Quit <- true
	}()
}

func (d *Dispatcher) Dispatch() {
	for {
		select {
		case job := <-JobQueue:
			go func(job Job) {
				jobChannel := <-d.WorkerPool
				jobChannel <- job
			}(job)
		case <-d.Quit:
			return
		}
	}
}

func entry(res http.ResponseWriter, req *http.Request) {
	// fetch job
	work := Job{serload: Serload{pri: "Just do it"}}
	JobQueue <- work
	fmt.Fprintf(res, "Hello World ...again")
}

func init() {
	runtime.GOMAXPROCS(MaxWorker)
	JobQueue = make(chan Job, MaxQueue)
	dispatcher := NewDispatcher(MaxWorker)
	dispatcher.Run()
}

// Go的协程和渠道
//这个demo的核心：
//1. 工作者工作协程，挂入调度器，取Job，执行Job，周而复始
//2. 调度器，从Job队列取Job，分配给工作者，周而复始
//3. web响应里，模拟了客户的请求-Job，并将此Job放入Job队列，只有有客户端请求，就周而复始的工作

func main5() {
	Port := "8086"
	IsHttp := true
	arg_num := len(os.Args)
	if 2 <= arg_num {
		Port = os.Args[1]
	}
	if 3 <= arg_num {
		if os.Args[2] == "true" {
			IsHttp = true
		} else {
			IsHttp = false
		}
	}
	fmt.Printf("server is http %t\n", IsHttp)
	fmt.Println("server listens at ", Port)

	http.HandleFunc("/", entry)

	var err error
	if IsHttp {
		err = http.ListenAndServe(":"+Port, nil)
	} else {
		err = http.ListenAndServeTLS(":"+Port, "server.crt", "server.key", nil)
	}
	if err != nil {
		fmt.Println("Server failure /// ", err)
	}

	fmt.Println("quit")
}
