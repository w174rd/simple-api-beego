package controllers

import (
	"encoding/json"
	"net/http"
	"regexp"
	"simple-api-beego/models"

	"github.com/beego/beego/orm"
	"github.com/beego/beego/v2/server/web"
)

type UserController struct {
	web.Controller
}

// Get all users
func (c *UserController) GetAll() {
	o := orm.NewOrm()
	var users []models.User
	_, err := o.QueryTable(new(models.User)).All(&users)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = map[string]string{"error": err.Error()}
	} else {
		c.Data["json"] = users
	}
	c.ServeJSON()
}

// Create a new user
func (c *UserController) Create() {
	var user models.User

	// Parse JSON dari request body
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &user)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]string{"error": "Invalid JSON format"}
		c.ServeJSON()
		return
	}

	// Validasi data user
	if validationErrors := validateUser(user); len(validationErrors) > 0 {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]interface{}{
			"error":   "Validation failed",
			"details": validationErrors,
		}
		c.ServeJSON()
		return
	}

	// Insert ke database
	o := orm.NewOrm()
	_, err = o.Insert(&user)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = map[string]string{"error": err.Error()}
	} else {
		c.Data["json"] = user
	}
	c.ServeJSON()
}

// Fungsi validasi data user
func validateUser(user models.User) map[string]string {
	errors := make(map[string]string)

	// Validasi nama
	if user.Name == "" {
		errors["name"] = "Name is required"
	}

	// Validasi email
	if user.Email == "" {
		errors["email"] = "Email is required"
	} else if !isValidEmail(user.Email) {
		errors["email"] = "Invalid email format"
	}

	return errors
}

// Helper untuk validasi email
func isValidEmail(email string) bool {
	// Regex sederhana untuk validasi email
	re := `^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`
	return regexp.MustCompile(re).MatchString(email)
}
