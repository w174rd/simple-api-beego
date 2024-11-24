package controllers

import (
	"encoding/json"
	"net/http"
	"simple-api-beego/models"
	"simple-api-beego/utils"

	"github.com/beego/beego/v2/server/web"
)

type AuthController struct {
	web.Controller
}

func (c *AuthController) Login() {
	var req models.User
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &req); err != nil {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]string{"error": "Invalid JSON format"}
		c.ServeJSON()
		return
	}

	user, err := GetUserByEmail(req.Email)
	if err != nil || user.Password != req.Password {
		// Ganti dengan pengecekan hash password jika menggunakan bcrypt
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]string{"error": "Invalid username or password"}
		c.ServeJSON()
		return
	}

	token, err := utils.GenerateToken(user.Id, user.Email)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]string{"error": "Failed to generate token"}
		c.ServeJSON()
		return
	}

	c.Data["json"] = map[string]string{"token": token}
	c.ServeJSON()
}
