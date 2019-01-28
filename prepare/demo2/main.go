package main

//执行shell命令 并且捕获输出
import (
	"fmt"
	"os/exec"
)

func main()  {
	var (
		cmd *exec.Cmd
		err error
		outPut []byte
	)

	cmd = exec.Command("/bin/bash","-c","ls -la;sleep 1 ;echo 2")

	if outPut, err = cmd.CombinedOutput();err != nil{
		fmt.Println(err)
		return
	}

	fmt.Println(string(outPut))

}
