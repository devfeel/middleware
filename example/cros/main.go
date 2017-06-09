package main

import (
	"fmt"
	"github.com/devfeel/dotweb"
	"github.com/devfeel/middleware/cros"
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
	_, err := ctx.WriteString("Index: ", dotweb.HeaderAccessControlAllowMethods, " => ", ctx.Response().QueryHeader(dotweb.HeaderAccessControlAllowMethods))
	return err
}

func Simple(ctx dotweb.Context) error {
	_, err := ctx.WriteString("Simple: ", dotweb.HeaderAccessControlAllowMethods, " => ", ctx.Response().QueryHeader(dotweb.HeaderAccessControlAllowMethods))
	return err
}

func NoCros(ctx dotweb.Context) error {
	_, err := ctx.WriteString("NoCros: ", dotweb.HeaderAccessControlAllowMethods, " => ", ctx.Response().QueryHeader(dotweb.HeaderAccessControlAllowMethods))
	return err
}

func InitRoute(server *dotweb.HttpServer) {
	server.Router().GET("/", Index).Use(NewCustomCROS())
	server.GET("/simple", Simple).Use(NewSimpleCROS())
	server.Router().GET("/nocros", NoCros)
}

// NewSimpleCROS
// default config:
// enabledCROS = true
// allowedOrigins = "*"
// allowedMethods = "GET, POST, PUT, DELETE, OPTIONS"
// allowedHeaders = "Content-Type"
// allowedP3P = "CP=\"CURa ADMa DEVa PSAo PSDo OUR BUS UNI PUR INT DEM STA PRE COM NAV OTC NOI DSP COR\""
func NewSimpleCROS() dotweb.Middleware {
	option := cros.NewConfig().UseDefault()
	return cros.Middleware(option)
}

func NewCustomCROS() dotweb.Middleware {
	option := cros.NewConfig()
	option.SetHeader("Content-Type")
	option.SetMethod("GET, POST")
	option.Enabled()
	return cros.Middleware(option)
}
