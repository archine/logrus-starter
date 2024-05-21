![](https://img.shields.io/badge/version-v1.x-green.svg) &nbsp; ![](https://img.shields.io/badge/version-go1.22-green.svg) &nbsp;  ![](https://img.shields.io/badge/builder-success-green.svg) &nbsp;


> Gin-plus快速集成Logrus日志   

## 一、安装
### 1. Get
```bash
go get github.com/archine/logrus-starter@v1.0.0
```
### 2. Mod
```bash
# go.mod文件加入下面的一条
github.com/archine/logrus-starter v1.0.0
# 命令行在该项目目录下执行
go mod tidy
```

## 二、使用说明
### 1. 配置
```yaml
log:
  level: debug # 日志级别 error、info、trace、warn、panic、fetal、debug
```
### 2. 替换框架内部日志
gin-plus框架中也需要打印部分日志，可以通过该方式替换默认的，可不设置
```go
package main

import (
	_ "gin-plus-demo/controller"
	"github.com/archine/gin-plus/v3/application"
	logrus_starter "github.com/archine/logrus-starter"
)

//go:generate gp-ast
func main() {
	application.Default().
		Log(&logrus_starter.Logger{}). // 替换默认日志
		Run()
}

```
### 3. 替换Gin默认的日志
由于 ``application.Default()`` 会采用Gin默认的日志，如果需要替换Gin默认的日志，可以通过``application.New()``方式来创建个干净的应用
```go
package main

import (
	_ "gin-plus-demo/controller"
	"github.com/archine/gin-plus/v3/application"
	logrus_starter "github.com/archine/logrus-starter"
)

//go:generate gp-ast
func main() {
	application.New(nil, logrus_starter.LogrusMiddleware()). // 替换默认日志
		Run()
}

```