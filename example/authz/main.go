package main

import (
	"fmt"
	"github.com/devfeel/dotweb"
	"github.com/devfeel/middleware/authz"
	"strconv"
	"github.com/casbin/casbin"
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

	//设置Casbin权限管理插件
	e := casbin.NewEnforcer("example/authz/authz_model.conf", "example/authz/authz_policy.csv")
	app.Use(authz.Middleware(e))

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

// Any handles all the test requests.
func Any(ctx dotweb.Context) error {
	return ctx.WriteString("Access allowed to: " + ctx.Request().Url())
}

// InitRoute initializes all the routes.
func InitRoute(server *dotweb.HttpServer) {
	server.Any("/", Any)
	server.Any("/dataset1/", Any)
	server.Any("/dataset2/", Any)
	server.Any("/dataset1/:name", Any)
	server.Any("/dataset2/:name", Any)
}
