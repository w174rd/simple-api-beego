package controllers

import (
	"encoding/json"
	"net/http"
	"simple-api-beego/models"
	"simple-api-beego/utils"

	"github.com/beego/beego/v2/server/web"
	"golang.org/x/crypto/bcrypt"
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
	// Verifikasi password
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]string{"error": "Invalid Email or password"}
		c.ServeJSON()
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]string{"error": "Invalid Email or password"}
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
