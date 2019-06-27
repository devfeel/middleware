package accesslog

import (
	"github.com/devfeel/dotweb"
	"net"
	"net/url"
	"strconv"
	"time"
	"unicode/utf8"
)

const logTarget = "dotweb_accesslog"

// Middleware new create a AccessLog Middleware
func Middleware() *AccessLogMiddleware {
	return &AccessLogMiddleware{}
}

// AccessLogMiddleware for logging HTTP requests in the Apache Common Log Format.
// Format: http://httpd.apache.org/docs/2.2/logs.html#common
// Write at 深圳华安大酒店
type AccessLogMiddleware struct {
	dotweb.BaseMiddlware
}

// Handle current middleware's handler
func (m *AccessLogMiddleware) Handle(ctx dotweb.Context) error {
	err := m.Next(ctx)
	log := buildContextLog(ctx)
	ctx.HttpServer().Logger().Raw(log, logTarget)
	return err
}

// buildContextLog builds a log entry for req in Apache Common Log Format.
func buildContextLog(ctx dotweb.Context) string {
	username := "-"
	if url.User != nil {
		if name := ctx.Request().URL.User.Username(); name != "" {
			username = name
		}
	}

	host, _, err := net.SplitHostPort(ctx.Request().RemoteAddr)

	if err != nil {
		host = ctx.Request().RemoteAddr
	}

	uri := ctx.Request().RequestURI

	// Requests using the CONNECT method over HTTP/2.0 must use
	// the authority field (aka r.Host) to identify the target.
	// Refer: https://httpwg.github.io/specs/rfc7540.html#CONNECT
	if ctx.Request().ProtoMajor == 2 && ctx.Request().Method == "CONNECT" {
		uri = ctx.Request().Host
	}
	if uri == "" {
		uri = ctx.Request().URL.RequestURI()
	}

	buf := make([]byte, 0, 3*(len(host)+len(username)+len(ctx.Request().Method)+len(uri)+len(ctx.Request().Proto)+50)/2)
	buf = append(buf, host...)
	buf = append(buf, " - "...)
	buf = append(buf, username...)
	buf = append(buf, " ["...)
	buf = append(buf, time.Now().Format("02/Jan/2006:15:04:05 -0700")...)
	buf = append(buf, `] "`...)
	buf = append(buf, ctx.Request().Method...)
	buf = append(buf, " "...)
	buf = appendQuoted(buf, uri)
	buf = append(buf, " "...)
	buf = append(buf, ctx.Request().Proto...)
	buf = append(buf, `" `...)
	buf = append(buf, strconv.Itoa(ctx.Response().Status)...)
	buf = append(buf, " "...)
	buf = append(buf, strconv.Itoa(int(ctx.Response().Size))...)
	return string(buf)
}

const lowerhex = "0123456789abcdef"

func appendQuoted(buf []byte, s string) []byte {
	var runeTmp [utf8.UTFMax]byte
	for width := 0; len(s) > 0; s = s[width:] {
		r := rune(s[0])
		width = 1
		if r >= utf8.RuneSelf {
			r, width = utf8.DecodeRuneInString(s)
		}
		if width == 1 && r == utf8.RuneError {
			buf = append(buf, `\x`...)
			buf = append(buf, lowerhex[s[0]>>4])
			buf = append(buf, lowerhex[s[0]&0xF])
			continue
		}
		if r == rune('"') || r == '\\' { // always backslashed
			buf = append(buf, '\\')
			buf = append(buf, byte(r))
			continue
		}
		if strconv.IsPrint(r) {
			n := utf8.EncodeRune(runeTmp[:], r)
			buf = append(buf, runeTmp[:n]...)
			continue
		}
		switch r {
		case '\a':
			buf = append(buf, `\a`...)
		case '\b':
			buf = append(buf, `\b`...)
		case '\f':
			buf = append(buf, `\f`...)
		case '\n':
			buf = append(buf, `\n`...)
		case '\r':
			buf = append(buf, `\r`...)
		case '\t':
			buf = append(buf, `\t`...)
		case '\v':
			buf = append(buf, `\v`...)
		default:
			switch {
			case r < ' ':
				buf = append(buf, `\x`...)
				buf = append(buf, lowerhex[s[0]>>4])
				buf = append(buf, lowerhex[s[0]&0xF])
			case r > utf8.MaxRune:
				r = 0xFFFD
				fallthrough
			case r < 0x10000:
				buf = append(buf, `\u`...)
				for s := 12; s >= 0; s -= 4 {
					buf = append(buf, lowerhex[r>>uint(s)&0xF])
				}
			default:
				buf = append(buf, `\U`...)
				for s := 28; s >= 0; s -= 4 {
					buf = append(buf, lowerhex[r>>uint(s)&0xF])
				}
			}
		}
	}
	return buf

}
