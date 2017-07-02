package basicauth

import (
	"github.com/devfeel/dotweb"
	"strings"
	"encoding/base64"
)
// basic auth option
type BasicAuthOption struct{
	//auth
	Auth func(name,pwd string) bool
}


// Middleware new create a AccessLog Middleware
func Middleware(option BasicAuthOption) *BasicAuthMiddleware {
	return &BasicAuthMiddleware{option:option}
}

//访问日志中间件
type BasicAuthMiddleware struct {
	dotweb.BaseMiddlware
	option BasicAuthOption
}

// Handle current middleware's handler
func (m *BasicAuthMiddleware) Handle(ctx dotweb.Context) error {
	name,pwd:=basicAuth(ctx)
	if name==""||pwd==""{
		return unAuthorized(ctx)
	}
	if m.option.Auth(name,pwd){
		m.Next(ctx)
	}else{
		return unAuthorized(ctx)
	}
	return nil
}
func unAuthorized(ctx dotweb.Context) error {
	ctx.Response().SetHeader("WWW-Authenticate","Basic realm=Input User&Password")
	ctx.Response().SetStatusCode(401)
	ctx.WriteString("Unauthorized")
	return nil
}
func basicAuth(ctx dotweb.Context) (string,string) {
	authorization:=ctx.Request().Header.Get("Authorization")
	if authorization!="" {
		authArr := strings.Split(authorization, " ")
		authDecode, _ := base64Decode([]byte(authArr[1]))
		authDecodeArr:=strings.Split(string(authDecode),":")
		return authDecodeArr[0],authDecodeArr[1]
	}
	return "",""
}

const (
	base64Table = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
)
var coder = base64.NewEncoding(base64Table)
func base64Encode(src []byte) []byte {
	return []byte(coder.EncodeToString(src))
}
func base64Decode(src []byte) ([]byte, error) {
	return coder.DecodeString(string(src))
}