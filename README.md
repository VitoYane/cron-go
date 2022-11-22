# cron-go
go版本的crontab, 兼容win, linux, darwin. Go练手小项目

Crontab 如果写成以下命令:
```
* * * * * cd d:\\ & echo %time% >> tmp.txt # 写入测试
```

使用本程序需要写成:
```
* * * * * !! cd d:\\ & echo %time% >> tmp.txt # 写入测试
```
中间加了2个感叹号, 主要为了分割调用频率和命令, 本可以写正则完全兼容, 但是go的cron库中, 还可以支持@hourly, @every 3s等写法, 所以使用分隔符

为了兼容多行, 中文, 符号等等, 带入参数使用base64编码.

使用方法:
```
go run cron.go KiAqICogKiAqICEhIGNkIGQ6XFwgJiBlY2hvICV0aW1lJSA+PiB0bXAudHh0ICMg5YaZ5YWl5rWL6K+V
```

带入后, 参数被解析为: * * * * * !! cd d:\\\\ & echo %time% >> tmp.txt # 写入测试

\# 可以作为注释内容
