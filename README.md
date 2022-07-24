<h1>
<a href="https://www.dtapp.net/">Golang Verify</a>
</h1>

📦 Golang 验证

[comment]: <> (go)
[![godoc](https://pkg.go.dev/badge/github.com/dtapps/goverify?status.svg)](https://pkg.go.dev/github.com/dtapps/goverify)
[![goproxy.cn](https://goproxy.cn/stats/github.com/dtapps/goverify/badges/download-count.svg)](https://goproxy.cn/stats/github.com/dtapps/goverify)
[![goreportcard.com](https://goreportcard.com/badge/github.com/dtapps/goverify    )](https://goreportcard.com/report/github.com/dtapps/goverify)
[![deps.dev](https://img.shields.io/badge/deps-go-red.svg)](https://deps.dev/go/github.com%2Fdtapps%2Fgoverify)

#### 安装

```go
go get -v -u github.com/dtapps/goverify
```

#### 使用

```go
package main

import (
	"github.com/dtapps/goverify"
	"testing"
)

// TestChinaMobile 验证手机号码
func TestChinaMobile(t *testing.T) {
	t.Log(goverify.ChinaMobile("13800138000"))
}

// TestChinaMobileNumber 验证中国移动手机号码
func TestChinaMobileNumber(t *testing.T) {
	t.Log(goverify.ChinaMobileNumber("13800138000"))
}

// TestChinaUnicomNumber 验证中国联通手机号码
func TestChinaUnicomNumber(t *testing.T) {
	t.Log(goverify.ChinaUnicomNumber("13800138000"))
}

// TestChinaTelecomNumber 验证中国电信手机号码
func TestChinaTelecomNumber(t *testing.T) {
	t.Log(goverify.ChinaTelecomNumber("13800138000"))
}
```

