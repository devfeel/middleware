# Domain
dotweb middleware for Domain permission.

## Install
```
go get -u github.com/devfeel/middleware
```

## Code：
```
import "github.com/devfeel/middleware/domain"

// config domain permission
domainConf := domain.NewAllowConfig()
domainConf.AddAllowDomain("www.dev.com")
domainConf.AddAllowDomain("127.0.0.1")
app.Use(domain.Middleware(domainConf))
```
## Config：

### SetMode

对config设置检查模式：
* OnlyAllow: only check allows domain permission
* OnlyReject: only check rejects domain permission

### AddAllowDomain

Add allow domain, only effective in OnlyAllow mode

### AddRejectDomain

Add reject domain, only effective in OnlyReject mode

#### NewAllowConfig

Return new domain config with OnlyAllow mode

#### NewRejectConfig

Return new domain config with OnlyReject mode

## Output

If no match config, will output like this:
* OnlyAllow: 403, not allow this domain visit this app
* OnlyReject: 403, this domain is rejected to visit this app
