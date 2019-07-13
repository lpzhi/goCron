package main

import (
	"fmt"
	"sync"
)

var jobChan chan  int
var  wg  sync.WaitGroup

func work(jobChan <- chan int)  {
	defer wg.Done()
	for job := range  jobChan  {
		fmt.Println("执行任务%d \n",job)
	}
}

func main()  {
	jobChan = make(chan int,100)

	//入队列

	for i:=0;i<100 ;i++  {
		jobChan <- i
	}
	wg.Add(1)
	close(jobChan)

	go work(jobChan)
	wg.Wait()
}
