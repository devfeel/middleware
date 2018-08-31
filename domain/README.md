# Domain
dotweb middleware for Domain permission.

## Install
```golang
go get -u github.com/devfeel/middleware
```

## Code：
```golang
import "github.com/devfeel/middleware/domain"

// config domain permission
domainConf := domain.NewAllowConfig()
domainConf.AddAllowDomain("www.dev.com")
domainConf.AddAllowDomain("127.0.0.1")
app.Use(domain.Middleware(domainConf))
```
## Config：

#### SetMode

对config设置检查模式：
* OnlyAllow: only check allows domain permission
* OnlyReject: only check rejects domain permission

#### AddAllowDomain

Add allow domain, only effective in OnlyAllow mode

#### AddRejectDomain

Add reject domain, only effective in OnlyReject mode

#### SetNotAllowHandle

Set handler which to be called when host not in allow domains

like:
~~~golang
func defaultNotAllowHandler(ctx dotweb.Context) error{
	return ctx.WriteStringC(http.StatusForbidden, NotAllowTip)
}

domainConf.SetNotAllowHandle(defaultNotAllowHandler)
~~~

#### SetRejectHandle

Set handler which to be called when host in reject domains

like:
~~~golang
func defaultRejectHandler(ctx dotweb.Context) error{
	return ctx.WriteStringC(http.StatusForbidden, RejectedTip)
}

domainConf.SetRejectHandle(defaultRejectHandler)
~~~

#### NewAllowConfig

Return new domain config with OnlyAllow mode

#### NewRejectConfig

Return new domain config with OnlyReject mode

## Output

If no match config, at default it will output like this:
* OnlyAllow: 403, not allow this domain visit this app
* OnlyReject: 403, this domain is rejected to visit this app
