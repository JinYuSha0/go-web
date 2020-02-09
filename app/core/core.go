package core

import (
	"github.com/kataras/iris"
	"gopkg.in/go-playground/validator.v9"
)

// 生成成功返回体
func GenSuccessRes(bo interface{}) iris.Map {
	return iris.Map{
		"status":  iris.StatusOK,
		"message": "ok",
		"bo":      bo,
	}
}

// 生成错误返回体
func GenErrorRes(status int, message string, bo interface{}) iris.Map {
	return iris.Map{
		"status":  status,
		"message": message,
		"bo":      bo,
	}
}

// 组合services
type Services func() (iris.Map, error)

func ServicesReduce(servs ...Services) func(iris.Context) (iris.Map, int, error) {
	return func(ctx iris.Context) (res iris.Map, statusCode int, err error) {
		statusCode = iris.StatusOK

		// 逐个service执行
		for _, f := range servs {
			res, err = f()

			// 遇到错误或者返回状态码非200 停止往下执行
			if err != nil || (res != nil && res["status"] != iris.StatusOK) {
				if err != nil {
					statusCode = iris.StatusInternalServerError
				} else {
					// iris中定义最大状态码为511所以规定自定义状态码必须大于511
					if res != nil && res["status"] != nil && res["status"].(int) < iris.StatusNetworkAuthenticationRequired {
						statusCode = res["status"].(int)
					}
				}

				break
			}
		}

		// 如果ctx传入非空则直接返回结果
		if ctx != nil && !ctx.IsStopped() {
			// 容错处理
			if res == nil {
				if err == nil {
					res = GenSuccessRes(nil)
				} else {
					res = GenErrorRes(statusCode, err.Error(), nil)
				}
			}

			ctx.StatusCode(statusCode)
			ctx.JSON(res)
		}

		return
	}
}

// 检查错误
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

// POST数据获取和校验
func BodyObtainAndValid(ctx iris.Context, obj interface{}, validate *validator.Validate) (err error) {
	if validate == nil {
		validate = validator.New()
	}

	// 获取数据
	if err = ctx.ReadJSON(obj); err != nil {
		panic(err)
	}

	// 执行校验
	if err = validate.Struct(obj); err != nil {
		if _, ok := err.(*validator.ValidationErrors); ok {
			panic(err)
		}

		if ctx != nil && !ctx.IsStopped() {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.JSON(GenErrorRes(iris.StatusBadRequest, err.Error(), nil))
		}
	}

	return
}

// 是否为空
func isNil(obj interface{}) bool {
	if obj == nil || obj == "" {
		return true
	}

	return false
}
