package helpers

import (
	"github.com/beego/beego/v2/server/web/context"
)

func Response(ctx *context.Context, code int, message string, data interface{}) {

	ctx.Output.Header("Content-Type", "application/json")

	ctx.Output.SetStatus(code)
	if code == 200 {
		ctx.Output.JSON(map[string]interface{}{
			"code":    code,
			"status":  "success",
			"message": message,
			"data":    data,
		}, false, false)
	} else {
		ctx.Output.JSON(map[string]interface{}{
			"code":    code,
			"status":  "error",
			"message": message,
		}, false, false)
	}

	// c.ServeJSON()
}
