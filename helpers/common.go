package helpers

import (
	"errors"
	"net/http"
	"reflect"
	"regexp"
	"strings"

	"github.com/beego/beego/v2/server/web/context"
)

// fungsi validasi
func ValidateRequiredFields(model interface{}, requiredFields []string) error {
	v := reflect.ValueOf(model).Elem()

	for _, field := range requiredFields {
		fieldValue := v.FieldByName(field)
		if !fieldValue.IsValid() || fieldValue.IsZero() {
			return errors.New(field + " is required")
		} else if strings.EqualFold("email", field) && !isValidEmail(fieldValue.Interface().(string)) {
			return errors.New("invalid email format")
		}
	}
	return nil
}

// Helper untuk validasi email
func isValidEmail(email string) bool {
	// Regex sederhana untuk validasi email
	re := `^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`
	return regexp.MustCompile(re).MatchString(email)
}

func ForbiddenHandler(ctx *context.Context) {
	ctx.Output.SetStatus(http.StatusForbidden) // Status 403 Forbidden
	ctx.Output.Body([]byte("403 Forbidden"))
}

func HandlePanic(ctx *context.Context) {
	if r := recover(); r != nil {
		// Jika ada panic, buat response JSON untuk error
		ctx.Output.SetStatus(500)
		ctx.Output.JSON(map[string]interface{}{
			"status":  "error",
			"message": "Internal Server Error",
			"error":   r,
		}, false, false)
	}
}
