package cronjob

import "github.com/rosefinchlj/easy-go/log"
import "github.com/robfig/cron/v3"

// cron job:
// 	- cron 表达式:
// 		- http://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/tutorial-lesson-06.html
// 		- https://en.wikipedia.org/wiki/Cron
//
// Cron 语法描述
// 每隔5秒执行一次：*/5 * * * * ?
// 每隔1分钟执行一次：0 */1 * * * ?
// 每天23点执行一次：0 0 23 * * ?
// 每天凌晨1点执行一次：0 0 1 * * ?
// 每月1号凌晨1点执行一次：0 0 1 1 * ?
// 在26分、29分、33分执行一次：0 26,29,33 * * * ?
// 每天的0点、13点、18点、21点都执行一次：0 0 0,13,18,21 * * ?

type CronJob struct {
    Cron *cron.Cron // 暴露原库 API
}

func New() *CronJob {
    return &CronJob{
        Cron: cron.New(
            cron.WithSeconds(), // 支持解析秒
            cron.WithChain(),
        ),
    }
}

// RegisterTask 注册 task:
func (m *CronJob) RegisterTask(tasks ...Task) (err error) {
    // batch register:
    for _, item := range tasks {
        // register:
        if entryID, err := m.Cron.AddFunc(item.Schedule, item.TaskFunc); err != nil {
            log.Errorf("cron job register tasks func error:, entryID=%v, err=%v", entryID, err)
        }
    }
    return err
}

// Run 注册和启动分开, 灵活调用位置
func (m *CronJob) Run(tasks ...Task) {
    m.RunAsync(tasks...)
}

// RunAsync 异步
func (m *CronJob) RunAsync(tasks ...Task) {
    // 允许在 run 中注册, 也可以分开, 传空即可
    _ = m.RegisterTask(tasks...)

    // 启动: 异步方式
    m.Cron.Start()
}

// RunSync 同步
func (m *CronJob) RunSync(tasks ...Task) {
    // 允许在 run 中注册, 也可以分开, 传空即可
    _ = m.RegisterTask(tasks...)

    // 启动: 同步方式
    m.Cron.Run()
}

// Stop 停止
func (m *CronJob) Stop() {
    m.Cron.Stop()
}
