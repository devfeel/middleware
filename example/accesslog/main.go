package main

import (
	"fmt"
	"github.com/devfeel/dotweb"
	"github.com/devfeel/middleware/accesslog"
	"strconv"
	"time"
)

func main() {
	//初始化DotServer
	app := dotweb.New()

	//设置dotserver日志目录
	//如果不设置，默认不启用，且默认为当前目录
	app.SetEnabledLog(true)
	app.SetLogPath("d:/gotmp/middleware/")

	//开启development模式
	app.SetDevelopmentMode()

	//设置路由
	InitRoute(app.HttpServer)

	//设置HttpModule
	//InitModule(app)

	//启动 监控服务
	//app.SetPProfConfig(true, 8081)

	// 开始服务
	port := 8080
	fmt.Println("dotweb.StartServer => " + strconv.Itoa(port))
	err := app.StartServer(port)
	fmt.Println("dotweb.StartServer error => ", err)
}

func Index(ctx dotweb.Context) error {
	time.Sleep(time.Millisecond * 10)
	_, err := ctx.WriteString("Index Use Accesslog")
	return err
}
func NoLog(ctx dotweb.Context) error {
	_, err := ctx.WriteString("NoLog")
	return err
}

func InitRoute(server *dotweb.HttpServer) {
	server.Router().GET("/", Index).Use(accesslog.Middleware())
	server.Router().GET("/nolog", NoLog)
}
