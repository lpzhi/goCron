package main

import (
	"fmt"
	"github.com/gorhill/cronexpr"
	"time"
)

//同时调度多个 cron表达式

type CronJob struct {
	expr     *cronexpr.Expression
	nextTime time.Time // expr.Next(now)
}

func main() {
	//需要有一个调度携程，他定时检查所有的任务，谁过期了就执行

	//1 定义两个job

	var (
		cronJob       *CronJob
		expr          *cronexpr.Expression
		now           time.Time
		scheduleTable map[string]*CronJob //key 任务的名字
	)

	scheduleTable = make(map[string]*CronJob)
	now = time.Now()
	expr = cronexpr.MustParse("*/5 * * * * * *")
	cronJob = &CronJob{
		expr:     expr,
		nextTime: expr.Next(now),
	}
	scheduleTable["job1"] = cronJob

	expr = cronexpr.MustParse("*/8 * * * * * *")
	cronJob = &CronJob{
		expr:     expr,
		nextTime: expr.Next(now),
	}
	scheduleTable["job2"] = cronJob

	////启动一个调度携程
	//go func() {
	//	//定时检查一下任务调度表
	//	for {
	//		for jb := range scheduleTable {
	//			fmt.Println(jb)
	//		}
	//	}
	//}()

	for {
		now = time.Now()
		for jb, v := range scheduleTable {
			//如果job时间 过期或者等于现在的时间
			if v.nextTime.Before(now) || v.nextTime.Equal(now) {
				go func() {
					fmt.Println(jb, v)
				}()
				v.nextTime = v.expr.Next(now)
			}

		}

		time.Sleep(time.Millisecond * 100)
	}

	time.Sleep(time.Second * 5)
}
