package task

import (
	"fmt"
	task "github.com/devfeel/dottask"
	"goCron/prepare/dottask/model"
	"time"
)

type UserTask struct {
}

func (usertask *UserTask) Installhour(cxt *task.TaskContext) error {

	fmt.Println(time.Now().String(), "=> Loop_T")
	//统计数据
	table := "log_install_20190724" //+ time.Now().Format("20060102")
	rs := model.HourInstallStatis(table)

	//model.Treating(rs, rs, rs, table)

	//for v, k := range rs {
	//	fmt.Println(v, k)
	//}

	fmt.Println(rs)
	//入库
	return nil
}
