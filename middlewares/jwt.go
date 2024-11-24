package middlewares

import (
	"net/http"
	"simple-api-beego/utils"
	"strings"

	"github.com/beego/beego/v2/server/web/context"
)

func JWTMiddleware(ctx *context.Context) {
	token := ctx.Input.Header("Authorization")
	if token == "" {
		ctx.Output.SetStatus(http.StatusUnauthorized)
		ctx.Output.JSON(map[string]interface{}{
			"code":    http.StatusUnauthorized,
			"error":   true,
			"message": "Authorization token is required",
		}, false, false)
		return
	}

	// check Bearer
	if !strings.HasPrefix(token, "Bearer ") {
		ctx.Output.SetStatus(http.StatusUnauthorized)
		ctx.Output.JSON(map[string]interface{}{
			"code":    http.StatusUnauthorized,
			"error":   true,
			"message": "Authorization header must start with Bearer",
		}, false, false)
		return
	}

	// hapus Bearer
	tokenString := strings.TrimPrefix(token, "Bearer ")

	_, err := utils.ValidateToken(tokenString)
	if err != nil {
		ctx.Output.SetStatus(http.StatusUnauthorized)
		ctx.Output.JSON(map[string]interface{}{
			"code":    http.StatusUnauthorized,
			"error":   true,
			"message": "The provided token is invalid or expired",
		}, false, false)
		return
	}

	// // Jika perlu, bisa simpan informasi user di context
	// ctx.Input.SetData("username", claims["username"])
}
