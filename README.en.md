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

### 📖 Language

[简体中文](./README.md) | English

### ⌛ Start

#### Project configuration

* 1：Configure `Global GOPATH` & `Project GOPATH`
* 2：Configure `Environment`
    * `GOPROXY=https://goproxy.cn,direct`
    * `GOFLAGS=-buildvcs=false`
    * `ENVCONFIG=dev`

#### Install dependencies

```
go mod tidy
```

#### Run

```
go run main.go
```

#### compile to an executable file

```
go build main.go
```

### 🐳 Docker

#### Compile the Golang code to Linux executable file

```
set GOOS=linux

set GOARCH=amd64

go build main.go
```

#### Docker Build

```
docker build -t note-macaron .
```

#### Docker Run

```
docker run -d -p 18083:18083 --name note-macaron note-macaron
```

### 📜 Licence

[MIT License](https://opensource.org/licenses/MIT) Copyright (c) 2022 周博义
