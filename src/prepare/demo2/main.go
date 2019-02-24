package main

import (
	"os/exec"
	"fmt"
)

func main()  {
	var(
		cmd *exec.Cmd
		output []byte
		err error
	)
	cmd = exec.Command("/bin/bash", "-c", "ls -l")
	if output,err = cmd.CombinedOutput();err!=nil{
		fmt.Println("error:",err)
		return
	}
	fmt.Println(string(output))
}
