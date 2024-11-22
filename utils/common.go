package utils

import (
	"log"
	"regexp"
	"simple-api-beego/models"
)

func validator(data interface{}, requiredParams []string) map[string]string {
	errors := make(map[string]string)

	var val []interface{}

	switch v := data.(type) {
	case models.User:
		val = append(val, v.Id)
		val = append(val, v.Name)
		val = append(val, v.Email)
		val = append(val, v.DeletedAt)
		val = append(val, v.CreatedAt)
		val = append(val, v.UpdateAt)
	default:
		errors["error"] = "Unsupported type"
	}

	// Loop melalui requiredParams dan periksa apakah ada yang kosong
	for _, param := range requiredParams {
		log.Println(param)
		// value, exists := val[param]
		// if !exists || strings.TrimSpace(value) == "" {
		// 	errors[param] = fmt.Sprintf("%s is required", param)
		// }
	}

	return errors

	// errors := make(map[string]string)

	// // Lakukan validasi manual berdasarkan jenis data
	// switch v := data.(type) {
	// case models.User:
	// 	if v.Name == "" {
	// 		errors["name"] = "Name is required"
	// 	}
	// 	if v.Email == "" {
	// 		errors["email"] = "Email is required"
	// 	} else if !isValidEmail(v.Email) {
	// 		errors["email"] = "Invalid email format"
	// 	}
	// default:
	// 	errors["error"] = "Unsupported type"
	// }

	// return errors
}

// Helper untuk validasi email
func isValidEmail(email string) bool {
	// Regex sederhana untuk validasi email
	re := `^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`
	return regexp.MustCompile(re).MatchString(email)
}
