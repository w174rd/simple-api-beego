package controllers

import (
	"encoding/json"
	"net/http"
	"simple-api-beego/helpers"
	"simple-api-beego/models"

	"github.com/beego/beego/v2/server/web"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	web.Controller
}

func (c *AuthController) Login() {
	var req models.User
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &req); err != nil {
		helpers.Response(c.Ctx, http.StatusBadRequest, "Invalid JSON format", nil)
		return
	}

	user, err := GetUserByEmail(req.Email)
	// Verifikasi password
	if err != nil {
		helpers.Response(c.Ctx, http.StatusBadRequest, "Invalid Email or password", nil)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		helpers.Response(c.Ctx, http.StatusBadRequest, "Invalid Email or password", nil)
		return
	}

	token, err := helpers.GenerateToken(user.Id, user.Email)
	if err != nil {
		helpers.Response(c.Ctx, http.StatusBadRequest, "Failed to generate token", nil)
		return
	}

	user.Token = token
	helpers.Response(c.Ctx, 200, "success", models.UserLogin(*user))
}
