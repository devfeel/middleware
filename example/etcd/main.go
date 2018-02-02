package main

import (
	"github.com/devfeel/dotweb"
	"fmt"
	"strconv"
	"github.com/devfeel/middleware/etcd"
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

	//设置HttpModule
	//InitModule(app)

	app.Use(etcd.Middleware(""))	//启用etcd中间件

	//启动 监控服务
	//app.SetPProfConfig(true, 8081)

	// 开始服务
	port := 8080
	fmt.Println("dotweb.StartServer => " + strconv.Itoa(port))
	err := app.StartServer(port)
	fmt.Println("dotweb.StartServer error => ", err)
}

func Index(ctx dotweb.Context) error {
	return ctx.WriteString("Index: ", dotweb.HeaderAccessControlAllowMethods, " => ", ctx.Response().QueryHeader(dotweb.HeaderAccessControlAllowMethods))
}

func List(ctx dotweb.Context) error {
	return nil
}

func Add(ctx dotweb.Context) error {
	return ctx.WriteString("Add: ", dotweb.HeaderAccessControlAllowMethods,
		" => ", ctx.RouterNode(),
		" => ", ctx.Response().QueryHeader(dotweb.HeaderAccessControlAllowMethods))
}

func NoCros(ctx dotweb.Context) error {
	return ctx.WriteString("NoCros: ", dotweb.HeaderAccessControlAllowMethods, " => ", ctx.Response().QueryHeader(dotweb.HeaderAccessControlAllowMethods))
}


func InitRoute(server *dotweb.HttpServer) {
	server.GET("/", Index).Use(etcd.Middleware(""))
	server.GET("/list", List).Use(etcd.Middleware(""))
	server.GET("/add", Add).Use(etcd.Middleware(""))
}


