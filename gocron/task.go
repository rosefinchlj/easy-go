package cronjob

type Tasks struct {
    Task []Task
}

type Task struct {
    Name     string
    Schedule string // 执行计划周期: cron 表达式
    TaskFunc func() // 任务方法
}
