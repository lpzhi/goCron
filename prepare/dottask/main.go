package main

import (
	"fmt"
	task "github.com/devfeel/dottask"
	"time"
)

var service *task.TaskService

func Job_Test(cxt *task.TaskContext) error {
	fmt.Println(time.Now().String(), "=> Loop_Test")
	fmt.Println(cxt.TaskData)
	//time.Sleep(time.Second * 3)
	return nil
}
func Loop_Test(ctx *task.TaskContext) error {
	fmt.Println(time.Now().String(), " => Loop_Test")
	fmt.Println(ctx)
	time.Sleep(time.Second * 3)
	return nil
}

func main() {
	//step 1 : init new task service

	service = task.StartNewService()

	_, err := service.CreateCronTask("testcron", true, "48-5 */2 * * * *", Job_Test, "test_aaa")
	if err != nil {
		fmt.Println("service.CreateCronTask error! => ", err.Error())
	}
	_, err = service.CreateLoopTask("testloop", true, 0, 1, Loop_Test, nil)
	if err != nil {
		fmt.Println("service.CreateLoopTask error! => ", err.Error())
	}
	service.StartAllTask()

	fmt.Println(service.PrintAllCronTask())

	for true {
		time.Sleep(time.Second * 10)
	}
}
