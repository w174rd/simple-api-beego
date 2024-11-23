package utils

import (
	"errors"
	"reflect"
	"regexp"
	"strings"
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
