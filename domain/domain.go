package domain

import (
	"github.com/devfeel/dotweb"
	"errors"
	"strings"
	"net/http"
)

var (
	NotAllowError = errors.New("not allow this domain visit this app")
	RejectedError = errors.New("this domain is rejected to visit this app")
	NotAllowTip = "not allow this domain visit this app"
	RejectedTip = "this domain is rejected to visit this app"
)

const(
	OnlyAllow = 1
	OnlyReject = 2
)

// DomainConfig the config used to set permission with domain
type DomainConfig struct{
	mode int
	allows map[string]struct{}
	rejects map[string]struct{}
}

// NewAllowConfig return new domain config with OnlyAllow mode
func NewAllowConfig() *DomainConfig{
	c := NewDomainConfig()
	c.mode = OnlyAllow
	return c
}

// NewRejectConfig return new domain config with OnlyReject mode
func NewRejectConfig() *DomainConfig{
	c := NewDomainConfig()
	c.mode = OnlyReject
	return c
}

// NewDomainConfig return new domain config
func NewDomainConfig() *DomainConfig{
	return &DomainConfig{
		mode : OnlyAllow,
		allows:make(map[string]struct{}),
		rejects:make(map[string]struct{}),
	}
}

// SetMode set check mode, now only support OnlyAllow or OnlyReject
func (c *DomainConfig) SetMode(mode int){
	if mode != OnlyAllow && mode != OnlyReject {
		c.mode = OnlyAllow
	}else{
		c.mode = mode
	}
}

// AddAllowDomain add allow domain, only effective in OnlyAllow mode
func (c *DomainConfig) AddAllowDomain(domain string){
	c.allows[domain] = struct{}{}
}

// AddRejectDomain add reject domain, only effective in OnlyReject mode
func (c *DomainConfig) AddRejectDomain(domain string){
	c.rejects[domain] = struct{}{}
}



// Middleware new create a AccessLog Middleware
func Middleware(conf *DomainConfig) *DomainMiddleware {
	return &DomainMiddleware{ config:conf}
}

// DomainMiddleware the middleware used to check domain visit permission
type DomainMiddleware struct {
	dotweb.BaseMiddlware
	config *DomainConfig
}

// Handle current middleware's handler
func (m *DomainMiddleware) Handle(ctx dotweb.Context) error {
	host := ctx.Request().Host
	domain := host[0:strings.Index(host, ":")]
	if m.config.mode == OnlyAllow{
		if !existsDomain(m.config.allows, domain){
			return ctx.WriteStringC(http.StatusForbidden, NotAllowTip)
		}
	}
	if m.config.mode == OnlyReject{
		if existsDomain(m.config.rejects, domain){
			return ctx.WriteStringC(http.StatusForbidden, RejectedTip)
		}
	}
	err := m.Next(ctx)
	return err
}


func existsDomain(m map[string]struct{}, domain string) bool{
	_, exists := m[domain]
	return exists
}