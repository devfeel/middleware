# Authz
dotweb middleware for authorization based on [Casbin](https://github.com/casbin/casbin).

## Install
```
go get -u github.com/devfeel/middleware
```

## Code：
```go
import "github.com/devfeel/middleware/cros"

// 设置Casbin Authz插件
e := casbin.NewEnforcer("path/to/model.conf", "path/to/policy.csv")
server.GET("/", Index).Use(authz.NewMiddleware(e))
```

## Documentation

The authorization determines a request based on ``{subject, object, action}``, which means what ``subject`` can perform what ``action`` on what ``object``. In this plugin, the meanings are:

1. ``subject``: the logged-on user name
2. ``object``: the URL path for the web resource like "dataset1/item1"
3. ``action``: HTTP method like GET, POST, PUT, DELETE, or the high-level actions you defined like "read-file", "write-blog"


For how to write authorization policy and other details, please refer to [the Casbin's documentation](https://github.com/casbin/casbin).

## Getting Help

- [Casbin](https://github.com/casbin/casbin)
