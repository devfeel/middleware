package accesslog

import (
	"github.com/devfeel/dotweb"
	"github.com/devfeel/dotweb/framework/convert"
	"github.com/devfeel/dotweb/logger"
	"time"
)

const logTarget = "dotweb_accesslog"

// Middleware new create a AccessLog Middleware
func Middleware() *AccessLogMiddleware {
	return &AccessLogMiddleware{}
}

//访问日志中间件
type AccessLogMiddleware struct {
	dotweb.BaseMiddlware
}

// Handle current middleware's handler
func (m *AccessLogMiddleware) Handle(ctx dotweb.Context) error {
	start := time.Now()
	m.Next(ctx)
	timetaken := int64(time.Now().Sub(start) / time.Millisecond)
	log := ctx.Request().Url() + " " + logContext(ctx, timetaken)
	logger.Logger().Log(log, logTarget, "debug")
	return nil
}

// logContext get default log string
// logformat: method userip proto status reqbytelen resbytelen timetaken
func logContext(ctx dotweb.Context, timetaken int64) string {
	var reqbytelen, resbytelen, method, proto, status, userip string
	if ctx != nil {
		reqbytelen = convert.Int642String(ctx.Request().ContentLength)
		resbytelen = convert.Int642String(ctx.Response().Size)
		method = ctx.Request().Method
		proto = ctx.Request().Proto
		status = convert.Int2String(ctx.Response().Status)
		userip = ctx.RemoteIP()
	}

	log := method + " "
	log += userip + " "
	log += proto + " "
	log += status + " "
	log += reqbytelen + " "
	log += resbytelen + " "
	log += convert.Int642String(timetaken)

	return log
}
