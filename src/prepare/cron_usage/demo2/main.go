package main

import (
	"github.com/gorhill/cronexpr"
	"time"
	"fmt"
)

//代表一个任务
type CronJob struct {
   expr *cronexpr.Expression
   nextTime time.Time
}

var(
    cronJob *CronJob
    expr *cronexpr.Expression
    now time.Time
    jobName string
    scheduleTable map[string]*CronJob
)

//调度多个cron任务
func main()  {
	now = time.Now()
	scheduleTable = make(map[string]*CronJob)
	expr = cronexpr.MustParse("*/5 * * * * * *")
	cronJob = &CronJob{
		expr:expr,
		nextTime:expr.Next(now),
	}
	scheduleTable["job1"] = cronJob


	expr = cronexpr.MustParse("*/5 * * * * * *")
	cronJob = &CronJob{
		expr:expr,
		nextTime:expr.Next(now),
	}
	scheduleTable["job2"] = cronJob

	go func() {
         for{
         	now = time.Now()
         	for jobName,cronJob = range scheduleTable{
         		if cronJob.nextTime.Before(now) || cronJob.nextTime.Equal(now){
         			go func(jobName string) {
         				fmt.Println(jobName,"执行成功......")
					}(jobName)
         			cronJob.nextTime = cronJob.expr.Next(now)
				}
			}

			 select {
			 case <-time.NewTimer(100 * time.Millisecond).C:
			}

		 }

	}()



time.Sleep(100 * time.Second)
}



