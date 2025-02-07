package application

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Application struct {
	Mongo struct {
		Host       string
		Port       string
		Username   string
		Password   string
		Database   string
		Collection string
	}
}

var App = &Application{}

func init() {
	// ioutil.ReadFile() 读取 yaml 配置文件
	file, err := ioutil.ReadFile("./application.yaml")
	if err != nil {
		log.Println(err)
	}

	// yaml.Unmarshal() 将 .yaml 配置文件中的配置解析到 Application 类型的变量中
	err = yaml.Unmarshal(file, App)
	if err != nil {
		log.Println(err)
	}
}
