package main

import (
	"github.com/gorhill/cronexpr"
	"fmt"
)

func main()  {
	var (
		cxpr *cronexpr.Expression
		err error
	)
	//分钟，小时，天，月，星期
if cxpr,err = cronexpr.Parse("* * * * *"); err!=nil{
    fmt.Println("error:",err)
    return
}

cxpr = cxpr

}


