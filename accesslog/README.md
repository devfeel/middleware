# AccessLog
dotweb middleware for logging HTTP requests in the Apache Common Log Format.

Log Format: http://httpd.apache.org/docs/2.2/logs.html#common

## Install
```
go get -u github.com/devfeel/middleware
```

## Code：
```
import "github.com/devfeel/middleware/accesslog"

server.GET("/", Index).Use(accesslog.Middleware())
```
## Config:

!!It's not suggest custom config.

default:
* use dotweb.logger
* use raw level
* logtarget is "dotweb_accesslog", like "dotweb_accesslog_2017_06_09.log"
* log-example: 127.0.0.1 - frank [10/Oct/2000:13:55:36 -0700] "GET /apache_pb.gif HTTP/1.0" 200 2326