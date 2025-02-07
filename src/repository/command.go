package repository

import (
	"go.mongodb.org/mongo-driver/bson"
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
