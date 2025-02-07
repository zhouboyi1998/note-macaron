package repository

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/macaron.v1"
	"log"
	"note-macaron/src/application"
	"note-macaron/src/model"
)

// Collection 连接 MongoDB, 连接指定的文档集合
func Collection(ctx *macaron.Context) *mongo.Collection {
	// 从配置文件中读取连接配置
	uri := "mongodb://" + application.App.Mongo.Username + ":" + application.App.Mongo.Password + "@" + application.App.Mongo.Host + ":" + application.App.Mongo.Port + "/"

	// 连接 MongoDB 数据库
	client, err := mongo.Connect(ctx.Req.Context(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Println(err)
	}

	// 连接配置文件指定的数据库和文档集合
	collection := client.Database(application.App.Mongo.Database).Collection(application.App.Mongo.Collection)

	return collection
}

// One 使用 ObjectId 查询单条命令
func One(ctx *macaron.Context) model.Command {
	// 获取集合连接
	collection := Collection(ctx)

	// 获取命令 Id
	commandId := ctx.Params(":commandId")
	// 将命令 Id 转换为 ObjectId 格式
	objectId, errHex := primitive.ObjectIDFromHex(commandId)
	if errHex != nil {
		log.Println(errHex)
	}

	// 使用 ObjectId 查询数据
	result := collection.FindOne(ctx.Req.Context(), bson.M{
		"_id": objectId,
	})

	// 命令结构体对象
	var command model.Command
	// 将查询结果赋值给命令结构体对象
	err := result.Decode(&command)
	if err != nil {
		log.Println(err)
	}

	return command
}

// OneByName 使用命令名称查询单条命令
func OneByName(ctx *macaron.Context) model.Command {
	// 获取集合连接
	collection := Collection(ctx)

	// 获取命令名称
	commandName := ctx.Params(":commandName")
	// 使用命令名称查询数据
	result := collection.FindOne(ctx.Req.Context(), bson.M{
		"command": commandName,
	})

	// 命令结构体对象
	var command model.Command
	// 将查询结果赋值给命令结构体对象
	err := result.Decode(&command)
	if err != nil {
		log.Println(err)
	}

	return command
}

// List 查询所有命令
func List(ctx *macaron.Context) []model.Command {
	// 获取集合连接
	collection := Collection(ctx)

	// 查询所有命令
	cursor, err := collection.Find(ctx.Req.Context(), bson.D{})
	if err != nil {
		log.Println(err)
	}

	// 命令结构体数组
	var commandArray []model.Command
	// 使用 cursor 指针遍历获取数据
	for cursor.Next(ctx.Req.Context()) {
		command := model.Command{}
		err := cursor.Decode(&command)
		if err != nil {
			log.Println(err)
		}
		commandArray = append(commandArray, command)
	}

	return commandArray
}

// ListByName 查询所有命令的名称
func ListByName(ctx *macaron.Context) []string {
	// 获取集合连接
	collection := Collection(ctx)

	// 查询所有命令
	cursor, err := collection.Find(ctx.Req.Context(), bson.D{})
	if err != nil {
		log.Println(err)
	}

	// 命令名称数组
	var nameArray []string
	// 使用 cursor 指针遍历获取数据
	for cursor.Next(ctx.Req.Context()) {
		command := model.Command{}
		err := cursor.Decode(&command)
		if err != nil {
			log.Println(err)
		}
		nameArray = append(nameArray, command.Command)
	}

	return nameArray
}

// InsertOne 新增单条命令
func InsertOne(ctx *macaron.Context) (*mongo.InsertOneResult, string) {
	// 获取集合连接
	collection := Collection(ctx)

	// 命令结构体对象
	var command model.Command
	// 将请求体参数赋值给命令结构体对象
	errBind := json.NewDecoder(ctx.Req.Request.Body).Decode(&command)
	if errBind != nil {
		log.Println(errBind)
	}
	// 生成 ObjectID
	command.Id = primitive.NewObjectID()

	// 新增单条命令
	result, err := collection.InsertOne(ctx.Req.Context(), command)
	if err != nil {
		log.Println(err)
	}

	return result, command.Command
}

// InsertMany 新增多条命令
func InsertMany(ctx *macaron.Context) *mongo.InsertManyResult {
	// 获取集合连接
	collection := Collection(ctx)

	// interface 数组
	var commandList []interface{}
	// 将请求体参数赋值给 interface 数组
	errBind := json.NewDecoder(ctx.Req.Request.Body).Decode(&commandList)
	if errBind != nil {
		log.Println(errBind)
	}

	// 新增多条命令
	result, err := collection.InsertMany(ctx.Req.Context(), commandList)
	if err != nil {
		log.Println(err)
	}

	return result
}

// UpdateOne 更新单条命令
func UpdateOne(ctx *macaron.Context) *mongo.UpdateResult {
	// 获取集合连接
	collection := Collection(ctx)

	// 命令结构体对象
	var command model.Command
	// 将请求体参数赋值给命令结构体对象
	errBind := json.NewDecoder(ctx.Req.Request.Body).Decode(&command)
	if errBind != nil {
		log.Println(errBind)
	}

	// 根据 ObjectId 更新单条命令
	result, err := collection.UpdateOne(ctx.Req.Context(), bson.M{"_id": command.Id}, bson.M{"$set": command})
	if err != nil {
		log.Println(err)
	}

	return result
}

// DeleteOne 删除单条命令
func DeleteOne(ctx *macaron.Context) (*mongo.DeleteResult, primitive.ObjectID) {
	// 获取集合连接
	collection := Collection(ctx)

	// 获取命令 Id
	commandId := ctx.Params(":commandId")
	// 将命令 Id 转换为 ObjectId 格式
	objectId, errHex := primitive.ObjectIDFromHex(commandId)
	if errHex != nil {
		log.Println(errHex)
	}

	// 根据 ObjectId 删除单条命令
	result, err := collection.DeleteOne(ctx.Req.Context(), bson.M{"_id": objectId})
	if err != nil {
		log.Println(err)
	}

	return result, objectId
}

// DeleteMany 删除多条命令
func DeleteMany(ctx *macaron.Context) (*mongo.DeleteResult, []primitive.ObjectID) {
	// 获取集合连接
	collection := Collection(ctx)

	// 命令 Id 数组
	var commandIds []string
	// 将请求体参数赋值给命令 Id 数组
	errBind := json.NewDecoder(ctx.Req.Request.Body).Decode(&commandIds)
	if errBind != nil {
		log.Println(errBind)
	}

	// ObjectId 数组
	var objectIds []primitive.ObjectID
	// 遍历命令 Id 数组, 转换为 ObjectId 格式
	for _, commandId := range commandIds {
		objectId, errHex := primitive.ObjectIDFromHex(commandId)
		if errHex != nil {
			log.Println(errHex)
		}
		objectIds = append(objectIds, objectId)
	}

	// 根据 ObjectId 数组删除多条命令
	result, err := collection.DeleteMany(ctx.Req.Context(), bson.M{"_id": bson.M{"$in": objectIds}})
	if err != nil {
		log.Println(err)
	}

	return result, objectIds
}
