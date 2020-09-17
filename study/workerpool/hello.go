package main

import (
	"fmt"
	"time"
)

var s = MyStruct{
	Name: "Chien",
	Age:  30,
}

type MyStruct struct {
	Name string
	Age  int
}

type WorkerPool struct {
	jobs    chan Job
	workers []Worker
}

func NewWorkerPool(workerNum int) (wp *WorkerPool) {
	wp = new(WorkerPool)
	wp.jobs = make(chan Job, 1000)
	for i := 0; i < workerNum; i++ {
		w := Worker{i}
		go w.run(wp.jobs)
	}
	return
}

func (wp WorkerPool) AddJob(job Job) {
	go func() { wp.jobs <- job }()
}

type Job struct {
	id int
}

type Worker struct {
	id int
}

func (w Worker) run(jobs chan Job) {
	fmt.Println("Worker", w.id, "on boarding")
	for job := range jobs {
		w.doJob(job)
	}
}

func (w Worker) doJob(job Job) {
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("Job [%d] processed by worker [%d]\n", job.id, w.id)
}

func main() {
	wp := NewWorkerPool(1000)
	for i := 0; i < 1000; i++ {
		wp.AddJob(Job{i})
	}
	time.Sleep(2 * time.Second)
}
