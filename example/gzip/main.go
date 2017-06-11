package main

import (
	"fmt"
	"github.com/devfeel/dotweb"
	"github.com/devfeel/middleware/gzip"
	"strconv"
)

func main() {
	//初始化DotServer
	app := dotweb.New()

	//设置dotserver日志目录
	//如果不设置，默认不启用，且默认为当前目录
	app.SetEnabledLog(true)

	//开启development模式
	app.SetDevelopmentMode()

	//设置路由
	InitRoute(app.HttpServer)

	// 开始服务
	port := 8080
	fmt.Println("dotweb.StartServer => " + strconv.Itoa(port))
	err := app.StartServer(port)
	fmt.Println("dotweb.StartServer error => ", err)
}

func Index(ctx dotweb.Context) error {
	ctx.Response().SetHeader(dotweb.HeaderContentType, dotweb.MIMETextPlainCharsetUTF8)
	ctx.WriteString(output)
	return nil
}

func NoUse(ctx dotweb.Context) error {
	ctx.WriteString(output)
	return nil
}

func InitRoute(server *dotweb.HttpServer) {
	server.Router().GET("/", Index).Use(gzip.Middleware(gzip.NewConfig().UseDefault()))
	server.Router().GET("/nouse", NoUse)
}

const output = `Features
支持静态路由、参数路由、组路由
路由支持文件/目录服务，支持设置是否允许目录浏览
中间件支持，支持App、Group、Router级别的设置 - https://github.com/devfeel/middleware
Feature支持，可绑定HttpServer全局启用
支持STRING/JSON/JSONP/HTML格式输出
统一的HTTP错误处理
统一的日志处理
支持Hijack与websocket
内建Cache支持
支持接入第三方模板引擎（需实现dotweb.Renderer接口）
模块可配置化，85%模块可通过配置维护
常规路由

支持GET\POST\HEAD\OPTIONS\PUT\PATCH\DELETE 这几类请求方法
支持HiJack\WebSocket\ServerFile三类特殊应用
支持Any注册方式，默认兼容GET\POST\HEAD\OPTIONS\PUT\PATCH\DELETE方式
支持通过配置开启默认添加HEAD方式
支持注册Handler，以启用配置化
支持检查请求与指定路由是否匹配
1、Router.GET(path string, handle HttpHandle)
2、Router.POST(path string, handle HttpHandle)
3、Router.HEAD(path string, handle HttpHandle)
4、Router.OPTIONS(path string, handle HttpHandle)
5、Router.PUT(path string, handle HttpHandle)
6、Router.PATCH(path string, handle HttpHandle)
7、Router.DELETE(path string, handle HttpHandle)
8、Router.HiJack(path string, handle HttpHandle)
9、Router.WebSocket(path string, handle HttpHandle)
10、Router.Any(path string, handle HttpHandle)
11、Router.RegisterRoute(routeMethod string, path string, handle HttpHandle)
12、Router.RegisterHandler(name string, handler HttpHandle)
13、Router.GetHandler(name string) (HttpHandle, bool)
14、Router.MatchPath(ctx Context, routePath string) bool
static router

静态路由语法就是没有任何参数变量，pattern是一个固定的字符串。

package main

import (
    "github.com/devfeel/dotweb"
)

func main() {
    dotapp := dotweb.New()
    dotapp.HttpServer.GET("/hello", func(ctx *dotweb.HttpContext) {
        ctx.WriteString("hello world!")
    })
    dotapp.StartServer(80)
}
parameter router

参数路由以冒号 : 后面跟一个字符串作为参数名称，可以通过 HttpContext的 GetRouterName 方法获取路由参数的值。

package main

import (
    "github.com/devfeel/dotweb"
)

func main() {
    dotapp := dotweb.New()
    dotapp.HttpServer.GET("/hello/:name", func(ctx dotweb.Context) error{
        _, err := ctx.WriteString("hello " + ctx.GetRouterName("name"))
        return err
    })
    dotapp.HttpServer.GET("/news/:category/:newsid", func(ctx dotweb.Context) error{
    	category := ctx.GetRouterName("category")
	    newsid := ctx.GetRouterName("newsid")
        _, err := ctx.WriteString("news info: category=" + category + " newsid=" + newsid)
        return err
    })
    dotapp.StartServer(80)
}
Middleware

支持粒度：App、Group、RouterNode
DotWeb.Use(m ...Middleware)
Group.Use(m ...Middleware)
RouterNode.Use(m ...Middleware)
启用顺序：App -> Group -> RouterNode，同级别下按Use的引入顺序执行
更多请参考：https://github.com/devfeel/middleware
JWT - example
AccessLog - example
CROS - example
Gzip
BasicAuth
Recover
HeaderOverride
 Server Config

HttpServer：

HttpServer.EnabledSession 设置是否开启Session支持，目前支持runtime、redis两种模式，默认不开启
HttpServer.EnabledGzip 设置是否开启Gzip支持，默认不开启
HttpServer.EnabledListDir 设置是否启用目录浏览，仅对Router.ServerFile有效，若设置该项，则可以浏览目录文件，默认不开启
HttpServer.EnabledAutoHEAD 设置是否自动启用Head路由，若设置该项，则会为除Websocket\HEAD外所有路由方式默认添加HEAD路由，默认不开启
Run Mode

新增development、production模式
默认development，通过DotWeb.SetDevelopmentMode\DotWeb.SetProductionMode开启相关模式
若设置development模式，未处理异常会输出异常详细信息
未来会拓展更多运行模式的配置
Exception

500 error

Default: 当发生未处理异常时，会根据RunMode向页面输出默认错误信息或者具体异常信息，并返回 500 错误头
User-defined: 通过DotServer.SetExceptionHandle(handler *ExceptionHandle)实现自定义异常处理逻辑
type ExceptionHandle func(Context, error)
404 error

Default: 当发生404异常时，会默认使用http.NotFound处理
User-defined: 通过DotWeb.SetNotFoundHandle(handler NotFoundHandle)实现自定义404处理逻辑
type NotFoundHandle  func(http.ResponseWriter, *http.Request)`
