# BasicAuth
dotweb middleware for BasicAuth.

## Install
```
go get -u github.com/devfeel/middleware
```

## Code：
```
import "github.com/devfeel/middleware/basicauth"

// 设置BasicAuth插件
option := basicauth.BasicAuthOption{}
	option.Auth= func(name, pwd string) bool {
		if name=="user"&& pwd=="123"{
			return true
		}
		return  false
	}
app.Use(basicauth.Middleware(option))
```

## Documentation

