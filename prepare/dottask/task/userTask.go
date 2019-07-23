package task

import (
	"fmt"
	task "github.com/devfeel/dottask"
	"time"
)

type UserTask struct {
}

func (usertask *UserTask) Installhour(cxt *task.TaskContext) error {

	fmt.Println(time.Now().String(), "=> Loop_Test")
	return nil
}
