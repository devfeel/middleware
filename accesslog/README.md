# AccessLog
dotweb middleware for AccessLog.

## Install
```
go get -u github.com/devfeel/middleware
```

## Code：
```
import "github.com/devfeel/middleware/accesslog"

// 设置jwt选项
server.GET("/", Index).Use(accesslog.Middleware())
```
## Config:

!!It's not suggest custom config.

default:
* use dotweb.logger
* use debug level
* logtarget is "dotweb_accesslog", like "dotweb_accesslog_debug_2017_06_09.log"
* logformat: method userip proto status reqbytelen resbytelen timetaken
* log-example: [debug] 2017-06-09 08:38:10.416369 [middleware.go:49] / GET 127.0.0.1 HTTP/1.1 200 0 19 10