package middlewares

import (
	"net/http"
	"simple-api-beego/helpers"
	"strings"

	"github.com/beego/beego/v2/server/web/context"
)

func JWTMiddleware(ctx *context.Context) {
	token := ctx.Input.Header("Authorization")

	if token == "" {
		helpers.Response(ctx, http.StatusUnauthorized, "Authorization token is required", nil)
		return
	}

	// check Bearer
	if !strings.HasPrefix(token, "Bearer ") {
		helpers.Response(ctx, http.StatusUnauthorized, "Authorization header must start with Bearer", nil)
		return
	}

	// hapus Bearer
	tokenString := strings.TrimPrefix(token, "Bearer ")

	_, err := helpers.ValidateToken(tokenString)
	if err != nil {
		helpers.Response(ctx, http.StatusUnauthorized, "The provided token is invalid or expired", nil)
		return
	}

	// // Jika perlu, bisa simpan informasi user di context
	// ctx.Input.SetData("username", claims["username"])
}
