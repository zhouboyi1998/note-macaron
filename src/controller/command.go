package controller

import (
	"encoding/json"
	"gopkg.in/macaron.v1"
	"note-macaron/src/repository"
)

// List 查询所有命令
func List(ctx *macaron.Context) {
	commandArray := repository.List(ctx)
	ctx.Resp.Header().Set("Content-Type", "application/json")
	json.NewEncoder(ctx.Resp).Encode(commandArray)
}
