package controllers

import (
	"fmt"
	"go-web/app/core"
	"go-web/app/services"
	"go-web/models"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate

type UserController struct {
	Ctx      iris.Context
	Service  services.UserService
	Validate *validator.Validate
}

func (c *UserController) BeforeActivation(b mvc.BeforeActivation) {
	// 中间件
	UserMiddleware := func(ctx iris.Context) {
		fmt.Println("user controller visit")
		ctx.Next()
	}
	b.Handle("POST", "/register", "PostRegister", UserMiddleware)
	b.Handle("POST", "/login", "PostLogin", UserMiddleware)
	b.Handle("GET", "/isExists", "GetIsExists", UserMiddleware)
}

func (c *UserController) PostRegister() {
	var user models.UserRegister

	// 获取并检验数据
	if err := core.BodyObtainAndValid(c.Ctx, &user, c.Validate); err != nil {
		return
	}

	register := func() (res iris.Map, err error) {
		res, err = c.Service.Register(user)
		return
	}

	core.ServicesReduce(register)(c.Ctx)
}

func (c *UserController) PostLogin() {
}

func (c *UserController) GetIsExists() {
	account := c.Ctx.URLParam("account")

	isExists := func() (res iris.Map, err error) {
		if err = c.Validate.Var(account, "required"); err != nil {
			res = core.GenErrorRes(iris.StatusBadRequest, err.Error(), nil)
			return res, nil
		}

		res, err = c.Service.IsExists(account)
		return
	}

	core.ServicesReduce(isExists)(c.Ctx)
}
