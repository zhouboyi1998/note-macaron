package router

import (
	"gopkg.in/macaron.v1"
	"note-macaron/src/controller"
)

// RegisterRouter 注册路由
func RegisterRouter(m *macaron.Macaron) {
	// 新建命令路由组
	m.Group("/command", func() {
		// 添加命令相关路由
		m.Get("/:commandId", controller.One)
		m.Get("/", controller.List)
		m.Post("/", controller.Insert)
		m.Post("/batch", controller.InsertBatch)
		m.Put("/", controller.Update)
		m.Put("/batch", controller.UpdateBatch)
		m.Delete("/:commandId", controller.Delete)
		m.Delete("/batch", controller.DeleteBatch)
		m.Get("/select/:commandName", controller.Select)
		m.Get("/name-list", controller.NameList)
	})
}
