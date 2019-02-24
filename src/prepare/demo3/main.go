package main

import (
	"os/exec"
	"context"
	"time"
)

//强制结束任务
func main()  {
	var (
		ctx context.Context
		cancelFunc context.CancelFunc
	)
	ctx,cancelFunc = context.WithCancel(context.TODO())
	go func() {
		exec.CommandContext(ctx,"/bin/bash","-c","sleep 2;echo hello")
	}()
	time.Sleep(1 * time.Second)
	//取消上下文
	cancelFunc()
}
