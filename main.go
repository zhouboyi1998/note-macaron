package main

import (
	"gopkg.in/macaron.v1"
	"note-macaron/src/application"
	"note-macaron/src/router"
)

func main() {
	// 新建 Macaron 实例
	app := macaron.Classic()
	// 注册路由
	router.RegisterRouter(app)
	// 启动服务
	app.Run(application.App.Server.Host, application.App.Server.Port)
}
