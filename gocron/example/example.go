package example

import (
    "github.com/rosefinchlj/easy-go/gocron"
    "github.com/rosefinchlj/easy-go/log"
    "time"
)

func Run() {
    cj := cronjob.New()

    // register + run:
    cj.Run(
        cronjob.Task{
            Name:     "Test 1",
            Schedule: "@every 1s",
            TaskFunc: func() {
                log.Infof("test1, every 1s, %v", time.Now())
            },
        },
        cronjob.Task{
            Name:     "Test 2",
            Schedule: "@every 2s",
            TaskFunc: func() {
                log.Infof("test2, every 2s, %v", time.Now())
            },
        },
    )

    // wait cron task:
    time.Sleep(time.Minute * 5)
}

func RegisterTask() {
    cj := cronjob.New()

    err := cj.RegisterTask(cronjob.Task{
        Name:     "Test 1",
        Schedule: "@every 1s",
        TaskFunc: func() {
            log.Infof("hello, every 1s, %v", time.Now())
        },
    })
    log.Infof("register err: %v", err)

    cj.Run()

    // wait cron task:
    time.Sleep(time.Minute * 5)
}
