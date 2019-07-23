package services

import (
	task "github.com/devfeel/dottask"
	task2 "goCron/prepare/dottask/task"
	"os"
	"runtime"
)

func RegisterTask(service *task.TaskService) {
	userTask := task2.UserTask{}
	service.RegisterHandler("Installhour", userTask.Installhour)
}

func LaodConfig(service *task.TaskService) {
	var configPath string
	currentPath := os.Getenv("DOTTASK_WORKER_DIR")

	if sysType := runtime.GOOS; sysType == "windows" {
		configPath = currentPath + "config\\task.xml.conf"
	} else {
		configPath = currentPath + "config/task.xml.conf"
	}

	service.LoadConfig(configPath)
}
