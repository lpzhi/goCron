package main

import (
	"fmt"
	"github.com/gorhill/cronexpr"
	"time"
)

func main()  {
	var(
		expr *cronexpr.Expression
		err error
		now time.Time
		nextTime time.Time
		ch chan int
	)
	ch = make(chan int)
	//那分钟（0-59）,那小时 (0-23) ，那天（1-31），哪月(1-12),星期几（0-6）
	//if expr,err = cronexpr.Parse("* * * * *");err!=nil {
	//	fmt.Println(err)
	//	return
	//}

	//每隔5分钟执行一次
	if expr,err = cronexpr.Parse("*/5 * * * * * *");err!=nil{
		fmt.Println(err)
		return
	}
	now = time.Now()
	//怎么计算下一次执行时间
	nextTime = expr.Next(now)
	//
	//fmt.Println(now,nextTime,nextTime.Sub(now),time.NewTimer(nextTime.Sub(now)))
	//
	//time.NewTimer(nextTime.Sub(now))

	time.AfterFunc(nextTime.Sub(now), func() {
		fmt.Println("被调度了:",nextTime)
		ch<-5
	})
	expr = expr


	fmt.Println(<-ch)
}
