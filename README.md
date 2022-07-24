<h1>
<a href="https://www.dtapp.net/">Golang Verify</a>
</h1>

📦 Golang 手机号验证

[comment]: <> (go)
[![godoc](https://pkg.go.dev/badge/github.com/dtapps/goverify?status.svg)](https://pkg.go.dev/github.com/dtapps/goverify)
[![goproxy.cn](https://goproxy.cn/stats/github.com/dtapps/goverify/badges/download-count.svg)](https://goproxy.cn/stats/github.com/dtapps/goverify)
[![goreportcard.com](https://goreportcard.com/badge/github.com/dtapps/goverify    )](https://goreportcard.com/report/github.com/dtapps/goverify)
[![deps.dev](https://img.shields.io/badge/deps-go-red.svg)](https://deps.dev/go/github.com%2Fdtapps%2Fgoverify)

#### 安装

```go
go get -v -u github.com/dtapps/goverify
```

#### 导入

```go
import "github.com/dtapps/goverify"
```

#### 使用

```go
package main

import (
	"github.com/dtapps/goverify"
	"testing"
)

func TestChinaMobile(t *testing.T) {
	t.Log(goverify.ChinaMobile("13800138000"))
}
```

