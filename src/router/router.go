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
		m.Get("/one/:commandId", controller.One)
		m.Get("/one/name/:commandName", controller.OneByName)
		m.Get("/list", controller.List)
		m.Get("/list/name", controller.ListByName)
		m.Post("/insert", controller.InsertOne)
		m.Post("/insert/many", controller.InsertMany)
		m.Put("/update", controller.UpdateOne)
		m.Delete("/delete/:commandId", controller.DeleteOne)
		m.Delete("/delete/many", controller.DeleteMany)
	})
}
