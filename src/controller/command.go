package controller

import (
	"encoding/json"
	"gopkg.in/macaron.v1"
	"note-macaron/src/repository"
)

// One 根据id查询命令
func One(ctx *macaron.Context) {
	command := repository.One(ctx)
	ctx.Resp.Header().Set("Content-Type", "application/json")
	json.NewEncoder(ctx.Resp).Encode(command)
}

// List 查询命令列表
func List(ctx *macaron.Context) {
	commandArray := repository.List(ctx)
	ctx.Resp.Header().Set("Content-Type", "application/json")
	json.NewEncoder(ctx.Resp).Encode(commandArray)
}

// Insert 新增命令
func Insert(ctx *macaron.Context) {
	result, commandName := repository.Insert(ctx)
	ctx.Resp.Header().Set("Content-Type", "application/json")
	json.NewEncoder(ctx.Resp).Encode(map[string]interface{}{
		"result":  result,
		"command": commandName,
	})
}

// InsertBatch 批量新增命令
func InsertBatch(ctx *macaron.Context) {
	result := repository.InsertBatch(ctx)
	ctx.Resp.Header().Set("Content-Type", "application/json")
	json.NewEncoder(ctx.Resp).Encode(result)
}

// Update 修改命令
func Update(ctx *macaron.Context) {
	result := repository.Update(ctx)
	ctx.Resp.Header().Set("Content-Type", "application/json")
	json.NewEncoder(ctx.Resp).Encode(result)
}

// UpdateBatch 批量修改命令
func UpdateBatch(ctx *macaron.Context) {
	result := repository.UpdateBatch(ctx)
	ctx.Resp.Header().Set("Content-Type", "application/json")
	json.NewEncoder(ctx.Resp).Encode(result)
}

// Delete 删除命令
func Delete(ctx *macaron.Context) {
	result, objectId := repository.Delete(ctx)
	ctx.Resp.Header().Set("Content-Type", "application/json")
	json.NewEncoder(ctx.Resp).Encode(map[string]interface{}{
		"result": result,
		"_id":    objectId,
	})
}

// DeleteBatch 批量删除命令
func DeleteBatch(ctx *macaron.Context) {
	result, objectIds := repository.DeleteBatch(ctx)
	ctx.Resp.Header().Set("Content-Type", "application/json")
	json.NewEncoder(ctx.Resp).Encode(map[string]interface{}{
		"result": result,
		"_ids":   objectIds,
	})
}

// Select 查询命令
func Select(ctx *macaron.Context) {
	command := repository.Select(ctx)
	ctx.Resp.Header().Set("Content-Type", "application/json")
	json.NewEncoder(ctx.Resp).Encode(command)
}

// NameList 查询命令名称列表
func NameList(ctx *macaron.Context) {
	nameArray := repository.NameList(ctx)
	ctx.Resp.Header().Set("Content-Type", "application/json")
	json.NewEncoder(ctx.Resp).Encode(nameArray)
}
