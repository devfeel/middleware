# Gzip
dotweb middleware for Gzip.

## Install
```
go get -u github.com/devfeel/middleware
```

## Code：
```
import "github.com/devfeel/middleware/gzip"

// 设置gzip选项
option := gzip.NewConfig().UseDefault()
server.GET("/", Index).Use(gzip.Middleware(option))
```
## Config：

#### UseDefault `func`

对config启用预设的默认值：
* GzipLevel = 1

### GzipLevel `int`

* 指定压缩等级，其值从1到9，1为最小化压缩（处理速度快），9为最大化压缩（处理速度慢）
* 默认为1

特别说明：
* 如果浏览器不支持gzip，无论是否启用该中间件，将不会启用gzip



