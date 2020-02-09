package middlewares

import (
	"fmt"
	"go-web/app/core"
	"runtime"

	"github.com/kataras/iris"
)

func CustomRecover(ctx iris.Context) {
	defer func() {
		if err := recover(); err != nil {
			if ctx.IsStopped() {
				return
			}

			var stacktrace string
			for i := 1; ; i++ {
				_, f, l, got := runtime.Caller(i)
				if !got {
					break
				}
				stacktrace += fmt.Sprintf("%s:%d\n", f, l)
			}

			errMsg := fmt.Sprintf("Error message: %s", err)
			logMessage := fmt.Sprintf("Error from: ('%s')\n", ctx.HandlerName())
			logMessage += errMsg + "\n"
			logMessage += fmt.Sprintf("%s\n", stacktrace)

			ctx.Application().Logger().Error(logMessage)

			ctx.StatusCode(500)
			ctx.JSON(core.GenErrorRes(iris.StatusInternalServerError, errMsg, nil))
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.StopExecution()
		}
	}()

	ctx.Next()
}
