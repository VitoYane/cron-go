# cron-go
go版本的crontab, 兼容win, linux, darwin. Go练手小项目

Crontab 如果写成以下命令:
```
* * * * * cd d:\\ & echo %time% >> tmp.txt # 写入测试
```

使用本程序需要写成:
```
* * * * * * !! cd d:\\ & echo %time% >> tmp.txt # 写入测试
```
- 中间加了2个感叹号, 主要为了分割调用频率和命令, 本可以写正则完全兼容, 但是go的cron库中, 还可以支持@hourly, @every 3s等写法, 所以使用分隔符
- 最前面多了一个星号, 是秒

为了兼容多行, 中文, 符号等等, 带入参数使用base64编码.

使用方法:
```
go run cron.go KiAqICogKiAqICogISEgY2QgYzogJiBlY2hvICV0aW1lJSA+PiB0bXAudHh0ICMg5YaZ5YWl5rWL6K+V
go run cron.go "* * * * * * !! cd c: & echo %time% >> tmp.txt # 写入测试"
```
两种写法都可以, 推荐使用base64, base64编码后带入, 避免shell处理特殊字符导致命令异常

\# 可以作为注释内容


## go 的库 cron 使用方法 Usage ¶
```
参考地址: https://pkg.go.dev/github.com/robfig/cron

Callers may register Funcs to be invoked on a given schedule. Cron will run them in their own goroutines.

c := cron.New()
c.AddFunc("0 30 * * * *", func() { fmt.Println("Every hour on the half hour") })
c.AddFunc("@hourly",      func() { fmt.Println("Every hour") })
c.AddFunc("@every 1h30m", func() { fmt.Println("Every hour thirty") })
c.Start()
..
// Funcs are invoked in their own goroutine, asynchronously.
...
// Funcs may also be added to a running Cron
c.AddFunc("@daily", func() { fmt.Println("Every day") })
..
// Inspect the cron job entries' next and previous run times.
inspect(c.Entries())
..
c.Stop()  // Stop the scheduler (does not stop any jobs already running).
CRON Expression Format ¶
A cron expression represents a set of times, using 6 space-separated fields.

Field name   | Mandatory? | Allowed values  | Allowed special characters
----------   | ---------- | --------------  | --------------------------
Seconds      | Yes        | 0-59            | * / , -
Minutes      | Yes        | 0-59            | * / , -
Hours        | Yes        | 0-23            | * / , -
Day of month | Yes        | 1-31            | * / , - ?
Month        | Yes        | 1-12 or JAN-DEC | * / , -
Day of week  | Yes        | 0-6 or SUN-SAT  | * / , - ?
Note: Month and Day-of-week field values are case insensitive. "SUN", "Sun", and "sun" are equally accepted.
```