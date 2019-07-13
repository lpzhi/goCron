package main

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

type result struct {
	err error
	outPut []byte
}

func main()  {
	var (
		ctx context.Context
		cancelFunc context.CancelFunc
		cmd *exec.Cmd
		resultChan chan *result
		res *result
	)

	resultChan = make(chan *result,1000)
	//执行一个cmd,让他在一个携程里面去执行，让它执行2秒
	//1秒的时候 我们杀死cmd
	ctx,cancelFunc = context.WithCancel(context.TODO())

	go func() {
		var (
			outPut []byte
			err error
		)

		cmd = exec.CommandContext(ctx,"/bin/bash","-c","sleep 2;echo hello")

		outPut,err = cmd.CombinedOutput()

		resultChan<-&result{
			err:err,
			outPut:outPut,
		}

	}()


	time.Sleep(1 * time.Second)
	cancelFunc()

	res = <- resultChan

	fmt.Println(res.err,res.outPut)
}



