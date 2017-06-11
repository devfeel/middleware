package gzip

import (
	"bufio"
	"compress/gzip"
	"errors"
	"github.com/devfeel/dotweb"
	"io"
	"net"
	"net/http"
	"strings"
)

const DefaultGzipLevel = 1
const gzipScheme = "gzip"

//Gzip配置
type (
	Config struct {
		GzipLevel int
	}

	gzipResponseWriter struct {
		io.Writer
		http.ResponseWriter
	}
)

/*gzipResponseWriter*/
func (w *gzipResponseWriter) WriteHeader(code int) {
	if code == http.StatusNoContent {
		w.ResponseWriter.Header().Del(dotweb.HeaderContentEncoding)
	}
	w.ResponseWriter.WriteHeader(code)
}

// Write writes the gzip data
func (w *gzipResponseWriter) Write(b []byte) (int, error) {
	if w.Header().Get(dotweb.HeaderContentType) == "" {
		w.Header().Set(dotweb.HeaderContentType, http.DetectContentType(b))
	}
	return w.Writer.Write(b)
}

// Flush implement http.Flusher interface
func (w *gzipResponseWriter) Flush() {
	w.Writer.(*gzip.Writer).Flush()
}

func (w *gzipResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	return w.ResponseWriter.(http.Hijacker).Hijack()
}

func (c *Config) UseDefault() *Config {
	c.GzipLevel = DefaultGzipLevel
	return c
}

func NewConfig() *Config {
	return &Config{}
}

// GzipMiddleware middleware with gzip
type GzipMiddleware struct {
	dotweb.BaseMiddlware
	config *Config
}

func (m *GzipMiddleware) Handle(ctx dotweb.Context) error {

	// check browser's Accept-Encoding contains "gzip"
	// if not support, no use gizp
	if strings.Contains(ctx.Request().QueryHeader(dotweb.HeaderAcceptEncoding), gzipScheme) {
		gw, err := gzip.NewWriterLevel(ctx.Response().Writer(), m.config.GzipLevel)
		if err != nil {
			return errors.New("use gzip error -> " + err.Error())
		}
		defer func() {
			gw.Close()
		}()

		grw := &gzipResponseWriter{Writer: gw, ResponseWriter: ctx.Response().Writer()}
		ctx.Response().SetWriter(grw)
		ctx.Response().SetHeader(dotweb.HeaderContentEncoding, gzipScheme)
		ctx.Response().SetHeader(dotweb.HeaderVary, dotweb.HeaderAcceptEncoding)
	}
	m.Next(ctx)

	return nil
}

// Middleware new create a Gzip Middleware
func Middleware(config *Config) *GzipMiddleware {
	return &GzipMiddleware{config: config}
}
