package main

import (
	"fmt"
	"github.com/devfeel/dotweb"
	"github.com/devfeel/middleware/cors"
)

func main() {
	//初始化DotServer
	app := dotweb.New()

	app.SetEnabledLog(true)

	//开启development模式
	app.SetDevelopmentMode()

	//设置路由
	InitRoute(app.HttpServer)

	// 开始服务
	port := 8080
	err := app.StartServer(port)
	if err != nil {
		fmt.Println("dotweb.StartServer error => ", err)
	}
}

func Index(ctx dotweb.Context) error {
	ctx.WriteString("Index: ", dotweb.HeaderAccessControlAllowMethods, " => ", ctx.Response().QueryHeader(dotweb.HeaderAccessControlAllowMethods))
	return nil
}

func Simple(ctx dotweb.Context) error {
	ctx.WriteString("Simple: ", dotweb.HeaderAccessControlAllowMethods,
		" => ", ctx.RouterNode(),
		" => ", ctx.Response().QueryHeader(dotweb.HeaderAccessControlAllowMethods))
	return nil
}

func NoCros(ctx dotweb.Context) error {
	ctx.WriteString("NoCros: ", dotweb.HeaderAccessControlAllowMethods, " => ", ctx.Response().QueryHeader(dotweb.HeaderAccessControlAllowMethods))
	return nil
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
	return cors.DefaultMiddleware()
}

func NewCustomCROS() dotweb.Middleware {
	option := cors.NewConfig()
	option.SetHeader("Content-Type")
	option.SetMethod("GET, POST")
	option.Enabled()
	return cors.Middleware(option)
}
