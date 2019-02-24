package main

import (
	"github.com/gorhill/cronexpr"
	"fmt"
	"time"
)

func main()  {
	var (
		cxpr *cronexpr.Expression
		err error
		now time.Time
		nextTime time.Time
	)
 //cron表达式:	//分钟，小时，天，月，星期
if cxpr,err = cronexpr.Parse("* * * * *"); err!=nil{
    fmt.Println("error:",err)
    return
}
//每隔5分钟执行一次
if cxpr,err = cronexpr.Parse("*/5 * * * *");err!=nil{
	fmt.Println("*/5 error:",err)
	return
}
//当前时间
now = time.Now()
//求下一次执行时间 
nextTime = cxpr.Next(now)
fmt.Println("当前时间 :",now,"--下一次执行时间为：",nextTime)

if cxpr,err = cronexpr.Parse("*/5 * * * * *");err!=nil{
	fmt.Println("*/5 errors:",err)
}
now = time.Now()
nextTime = cxpr.Next(now)
//定时期调度
time.AfterFunc(nextTime.Sub(now), func() {
	fmt.Println("被调度了")
})
time.Sleep( 10 * time.Second)
cxpr = cxpr

}


