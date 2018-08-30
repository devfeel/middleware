package main

import (
	"fmt"
	"github.com/devfeel/dotweb"
	"strconv"
	"github.com/devfeel/middleware/domain"
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

	/*domainConf := domain.NewAllowConfig()
	domainConf.AddAllowDomain("www.dev.com")
	domainConf.AddAllowDomain("127.0.0.1")
	app.Use(domain.Middleware(domainConf))*/

	domainConf := domain.NewRejectConfig()
	domainConf.AddRejectDomain("www.dev.com")
	app.Use(domain.Middleware(domainConf))

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
	return ctx.WriteString("Index")
}


func InitRoute(server *dotweb.HttpServer) {
	server.GET("/", Index)
}
