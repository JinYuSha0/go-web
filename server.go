package main

import (
	"go-web/app/controllers"
	"go-web/app/middlewares"
	"go-web/app/services"
	"go-web/dao"
	"go-web/utils"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"gopkg.in/go-playground/validator.v9"
)

func main() {
	// 获取配置
	config := utils.GetConfig()

	// 连接数据库
	dao.Connect()

	app := iris.New()

	// 异常捕获中间件
	app.Use(middlewares.CustomRecover)

	user := mvc.New(app.Party("/user"))
	userDAO := dao.NewUserDAO()
	user.Register(services.NewUserService(userDAO))
	user.Register(validator.New())
	user.Handle(new(controllers.UserController))

	app.Run(iris.Addr(":8080"), iris.WithConfiguration(config.Iris))
}
