package main

import (
	"fmt"
	"github.com/devfeel/dotweb"
	"github.com/devfeel/middleware/basicauth"
	"strconv"
	"time"
)

func main() {
	//初始化DotServer
	app := dotweb.New()

	//设置dotserver日志目录
	//如果不设置，默认不启用，且默认为当前目录
	app.SetEnabledLog(true)
	app.SetLogPath("f:/gotmp/middleware/")

	//开启development模式
	app.SetDevelopmentMode()

	//设置路由
	InitRoute(app.HttpServer)

	//设置HttpModule
	//InitModule(app)

	//启动 监控服务
	//app.SetPProfConfig(true, 8081)

	app.Use(NewBasicAuth(app))
	// 开始服务
	port := 8080
	fmt.Println("dotweb.StartServer => " + strconv.Itoa(port))
	err := app.StartServer(port)
	fmt.Println("dotweb.StartServer error => ", err)
}

func Index(ctx dotweb.Context) error {
	time.Sleep(time.Millisecond * 10)
	return ctx.WriteString("Authorization")
}
func Index1(ctx dotweb.Context) error {
	time.Sleep(time.Millisecond * 10)
	return ctx.WriteString("index1")
}

func InitRoute(server *dotweb.HttpServer) {
	server.GET("/", Index)
	server.GET("/index1", Index1)
}

func NewBasicAuth(app *dotweb.DotWeb) dotweb.Middleware {
	option := basicauth.BasicAuthOption{}
	option.Auth = func(name, pwd string) bool {
		if name == "user" && pwd == "123" {
			return true
		}
		return false
	}
	return basicauth.Middleware(option)
}
