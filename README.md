<h1 align="center">📔 note-macaron</h1>

<p align="center">
<a target="_blank" href="https://github.com/zhouboyi1998/note-macaron"> 
<img src="https://img.shields.io/github/stars/zhouboyi1998/note-macaron?logo=github">
</a>
<a target="_blank" href="https://opensource.org/licenses/MIT"> 
<img src="https://img.shields.io/badge/license-MIT-red"> 
</a>
<img src="https://img.shields.io/badge/Go-1.23-darkturquoise">
<img src="https://img.shields.io/badge/Macaron-1.5.0-steelblue">
<img src="https://img.shields.io/badge/MongoDB Go Driver-1.17.2-seagreen">
</p>

### 📖 语言

简体中文 | [English](./README.en.md)

### ⌛ 开始

#### 项目配置

* 1：配置 `Global GOPATH` & `Project GOPATH`
* 2：配置 `Environment`
    * `GOPROXY=https://goproxy.cn,direct`
    * `GOFLAGS=-buildvcs=false`
    * `ENVCONFIG=dev`

#### 安装依赖

```
go mod tidy
```

#### 运行

```
go run main.go
```

#### 编译成可执行文件

```
go build main.go
```

### 🐳 Docker

#### Go 代码编译成 Linux 可执行文件

```
set GOOS=linux

set GOARCH=amd64

go build main.go
```

#### Docker 构建

```
docker build -t note-macaron .
```

#### Docker 运行

```
docker run -d -p 18083:18083 --name note-macaron note-macaron
```

### 📜 开源协议

[MIT License](https://opensource.org/licenses/MIT) Copyright (c) 2022 周博义
