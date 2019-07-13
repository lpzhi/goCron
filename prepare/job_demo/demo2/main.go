package main

import (
	"fmt"
	"time"
)

func main()  {
	JobChan :=make(chan Job)

	go Work(JobChan)

	JobChan <- Job{}

	time.Sleep(time.Second * 1)
}

type Job struct {

}

func Work(jobChan <- chan Job )  {
	for job :=range jobChan {
		//process job
		fmt.Println("job",job)
	}
}

func TryEnqueue(job Job, jobChan chan<- Job) bool {
	select {
	case jobChan <- job:
		return true
	default:
		return false
	}
}