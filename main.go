package main

import (
	"encoding/json"
	"gopkg.in/macaron.v1"
	"net/http"
)

func main() {
	// 新建 Macaron 实例
	app := macaron.Classic()

	// Hello World
	app.Get("/hello/:name", func(ctx *macaron.Context, w http.ResponseWriter) {
		name := ctx.Params(":name")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"code":    http.StatusOK,
			"message": "Hello, " + name,
		})
	})

	// 启动服务
	app.Run(18083)
}
