package main

import (
	"fmt"
	task "github.com/devfeel/dottask"
	"goCron/prepare/dottask/services"
	"os"
	"os/signal"
	"syscall"
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

func startTask() {
	service = task.StartNewService()
	services.RegisterTask(service) //注册任务
	services.LaodConfig(service)
	service.StartAllTask()
	fmt.Println(service.PrintAllCronTask())
}

func main() {

	c := make(chan os.Signal)
	//signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		for s := range c {
			switch s {
			case syscall.SIGHUP, syscall.SIGTERM, syscall.SIGQUIT:
				fmt.Println("退出", s)
			case syscall.SIGINT:
				fmt.Println("退出重启", s)
				service.StopAllTask()
				startTask()
			default:
				fmt.Println("other", s)
			}
		}
	}()

	fmt.Println("进程启动...")
	startTask()

	for true {
		time.Sleep(time.Second * 1)
	}
}
