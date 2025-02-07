package router

import (
	"gopkg.in/macaron.v1"
	"note-macaron/src/controller"
)

// RegisterRouter 注册路由
func RegisterRouter(m *macaron.Macaron) {
	// 新建路由组
	m.Group("/note/command", func() {
		// 添加路由规则
		m.Get("/list", controller.List)
	})
}
