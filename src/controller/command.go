package controller

import (
	"encoding/json"
	"gopkg.in/macaron.v1"
	"note-macaron/src/repository"
)

// One 使用 ObjectId 查询单条命令
func One(ctx *macaron.Context) {
	command := repository.One(ctx)
	ctx.Resp.Header().Set("Content-Type", "application/json")
	json.NewEncoder(ctx.Resp).Encode(command)
}

// OneByName 使用命令名称查询单条命令
func OneByName(ctx *macaron.Context) {
	command := repository.OneByName(ctx)
	ctx.Resp.Header().Set("Content-Type", "application/json")
	json.NewEncoder(ctx.Resp).Encode(command)
}

// List 查询所有命令
func List(ctx *macaron.Context) {
	commandArray := repository.List(ctx)
	ctx.Resp.Header().Set("Content-Type", "application/json")
	json.NewEncoder(ctx.Resp).Encode(commandArray)
}

// ListName 查询所有命令的名称
func ListName(ctx *macaron.Context) {
	nameArray := repository.ListName(ctx)
	ctx.Resp.Header().Set("Content-Type", "application/json")
	json.NewEncoder(ctx.Resp).Encode(nameArray)
}

// InsertOne 新增单条命令
func InsertOne(ctx *macaron.Context) {
	result, commandName := repository.InsertOne(ctx)
	ctx.Resp.Header().Set("Content-Type", "application/json")
	json.NewEncoder(ctx.Resp).Encode(map[string]interface{}{
		"result":  result,
		"command": commandName,
	})
}

// InsertMany 新增多条命令
func InsertMany(ctx *macaron.Context) {
	result := repository.InsertMany(ctx)
	ctx.Resp.Header().Set("Content-Type", "application/json")
	json.NewEncoder(ctx.Resp).Encode(result)
}

// UpdateOne 更新单条命令
func UpdateOne(ctx *macaron.Context) {
	result := repository.UpdateOne(ctx)
	ctx.Resp.Header().Set("Content-Type", "application/json")
	json.NewEncoder(ctx.Resp).Encode(result)
}

// DeleteOne 删除单条命令
func DeleteOne(ctx *macaron.Context) {
	result, objectId := repository.DeleteOne(ctx)
	ctx.Resp.Header().Set("Content-Type", "application/json")
	json.NewEncoder(ctx.Resp).Encode(map[string]interface{}{
		"result": result,
		"_id":    objectId,
	})
}

// DeleteMany 删除多条命令
func DeleteMany(ctx *macaron.Context) {
	result, objectIds := repository.DeleteMany(ctx)
	ctx.Resp.Header().Set("Content-Type", "application/json")
	json.NewEncoder(ctx.Resp).Encode(map[string]interface{}{
		"result": result,
		"_ids":   objectIds,
	})
}
