package helpers

import (
	"net/http"

	"github.com/beego/beego/v2/server/web/context"
)

// fungsi untuk mengijinkan diakses sumber asal
func CORS(ctx *context.Context) {
	ctx.Output.Header("Access-Control-Allow-Origin", "*") // Izinkan semua asal
	ctx.Output.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	ctx.Output.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")

	// Tangani permintaan OPTIONS
	if ctx.Input.Method() == http.MethodOptions {
		ctx.ResponseWriter.WriteHeader(http.StatusNoContent)
		return
	}
}
