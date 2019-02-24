package main

import (
	"os/exec"
	"context"
	"time"
	"fmt"
)
type result struct {
	err error
	output string
}

//强制结束任务
func main()  {
	var (
		ctx context.Context
		cancelFunc context.CancelFunc
		resultChan chan *result
		res *result
		cmd *exec.Cmd
		output []byte
		err error
	)
	resultChan = make(chan *result,100)

	ctx,cancelFunc = context.WithCancel(context.TODO())
	go func() {
		cmd = exec.CommandContext(ctx,"/bin/bash","-c","sleep 2;echo hello")
		output,err = cmd.CombinedOutput()
        resultChan <- &result{
        	output:string(output),
        	err:err,
		}
	}()
	time.Sleep(1 * time.Second)
	//取消上下文
	cancelFunc()

	res = <- resultChan
	fmt.Println("output",res.output,"error:",res.err)
}
