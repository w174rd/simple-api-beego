package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"simple-api-beego/helpers"
	"simple-api-beego/models"
	"time"

	"github.com/beego/beego/orm"
	"github.com/beego/beego/v2/server/web"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	web.Controller
}

// Get all users
func (c *UserController) GetAll() {
	o := orm.NewOrm()
	var users []models.User
	_, err := o.QueryTable(new(models.User)).Filter("DeletedAt__isnull", true).All(&users)
	if err != nil {
		helpers.Response(c.Ctx, http.StatusInternalServerError, err.Error(), nil)
	} else {
		var newUsers []models.User
		for _, user := range users {
			newUsers = append(newUsers, models.UserDefault(user))
		}
		helpers.Response(c.Ctx, 200, "success", newUsers)
	}
}

// Fungsi untuk mendapatkan user berdasarkan ID
func (c *UserController) GetUserByID() {
	// Ambil ID dari parameter URL
	id, err := c.GetInt(":id")
	if err != nil {
		helpers.Response(c.Ctx, http.StatusBadRequest, "Invalid user ID", nil)
		return
	}

	// Ambil data user berdasarkan ID
	user, err := GetUserByID(id)
	if err != nil {
		helpers.Response(c.Ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	// Response user dalam format JSON
	helpers.Response(c.Ctx, 200, "success", models.UserDefault(*user))
}

// Create a new user
func (c *UserController) Create() {
	var user models.User

	// Parse JSON dari request body
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &user)
	if err != nil {
		helpers.Response(c.Ctx, http.StatusBadRequest, "Invalid JSON format", nil)
		return
	}

	// Validasi data user
	if err := helpers.ValidateRequiredFields(&user, []string{"Name", "Email", "Password"}); err != nil {
		helpers.Response(c.Ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	// Check existing user
	if _, err := GetUserByEmail(user.Email); err == nil {
		helpers.Response(c.Ctx, http.StatusBadRequest, "email has been registered", nil)
		return
	}

	// Enkripsi password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		helpers.Response(c.Ctx, http.StatusBadRequest, "Failed to hash password", nil)
		return
	}

	user.Password = string(hashedPassword)

	// Insert ke database
	o := orm.NewOrm()
	_, err = o.Insert(&user)
	if err != nil {
		helpers.Response(c.Ctx, http.StatusInternalServerError, err.Error(), nil)
	} else {
		helpers.Response(c.Ctx, 200, "success", models.UserComplete(user))
	}
}

func (c *UserController) Update() {
	// Ambil ID dari parameter URL
	id, errId := c.GetInt(":id")
	if errId != nil {
		helpers.Response(c.Ctx, http.StatusBadRequest, "Invalid user ID", nil)
		return
	}

	var user models.User
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &user)
	if err != nil {
		helpers.Response(c.Ctx, http.StatusBadRequest, "Invalid JSON format", nil)
		return
	}

	if err := helpers.ValidateRequiredFields(&user, []string{"Name", "Email"}); err != nil {
		helpers.Response(c.Ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	// Cek apakah user ada
	newUserData, err := GetUserByID(id)
	if err != nil {
		helpers.Response(c.Ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	newUserData.Name = user.Name
	newUserData.Email = user.Email

	user = *newUserData

	o := orm.NewOrm()
	if _, err := o.Update(&user); err != nil {
		helpers.Response(c.Ctx, http.StatusBadRequest, "Failed to update user", nil)
		return
	}

	// Response sukses
	helpers.Response(c.Ctx, 200, "success", models.UserDefault(user))
}

// Fungsi untuk menghapus user
func (c *UserController) Delete() {
	// Ambil ID dari parameter URL
	id, err := c.GetInt(":id")
	if err != nil {
		helpers.Response(c.Ctx, http.StatusBadRequest, "Invalid user ID", nil)
		return
	}

	// ORM instance
	o := orm.NewOrm()
	user := models.User{Id: id}

	// Periksa apakah user ada
	if err := o.Read(&user); err != nil {
		helpers.Response(c.Ctx, http.StatusBadRequest, "User not found", nil)
		return
	}

	// Set DeletedAt dengan waktu saat ini
	now := time.Now()
	user.DeletedAt = &now

	// Update user dengan DeletedAt yang telah diubah
	if _, err := o.Update(&user, "DeletedAt"); err != nil {
		helpers.Response(c.Ctx, http.StatusInternalServerError, "Failed to delete user", nil)
		return
	}

	// Response sukses
	helpers.Response(c.Ctx, 200, "User deleted successfully", nil)
}

func GetUserByID(id int) (*models.User, error) {
	// ORM instance
	o := orm.NewOrm()
	var user models.User
	err := o.QueryTable(new(models.User)).Filter("Id", id).Filter("DeletedAt__isnull", true).One(&user)
	if err != nil {
		if err == orm.ErrNoRows {
			return nil, errors.New("user not found")
		} else {
			return nil, errors.New("failed to retrieve user")
		}
	}

	return &user, nil
}

func GetUserByEmail(email string) (*models.User, error) {
	// ORM instance
	o := orm.NewOrm()
	var user models.User
	err := o.QueryTable(new(models.User)).Filter("Email", email).Filter("DeletedAt__isnull", true).One(&user)
	if err != nil {
		if err == orm.ErrNoRows {
			return nil, errors.New("user not found")
		} else {
			return nil, errors.New("failed to retrieve user")
		}
	}

	return &user, nil
}
