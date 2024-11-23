package controllers

import (
	"encoding/json"
	"net/http"
	"simple-api-beego/models"
	"simple-api-beego/utils"
	"time"

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
	_, err := o.QueryTable(new(models.User)).Filter("DeletedAt__isnull", true).All(&users)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = map[string]string{"error": err.Error()}
	} else {
		var newUsers []models.User
		for _, user := range users {
			newUsers = append(newUsers, models.UserDefault(user))
		}
		c.Data["json"] = newUsers
	}
	c.ServeJSON()
}

// Fungsi untuk mendapatkan user berdasarkan ID
func (c *UserController) GetUserByID() {
	// Ambil ID dari parameter URL
	id, err := c.GetInt(":id")
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]string{"error": "Invalid user ID"}
		c.ServeJSON()
		return
	}

	// ORM instance
	o := orm.NewOrm()
	user := models.User{Id: id}

	// Ambil data user berdasarkan ID
	err = o.QueryTable(new(models.User)).Filter("id", id).Filter("DeletedAt__isnull", true).One(&user)
	if err != nil {
		if err == orm.ErrNoRows {
			// Tidak ditemukan user dengan ID tersebut
			c.Ctx.Output.SetStatus(http.StatusBadRequest)
			c.Data["json"] = map[string]string{"error": "User not found"}
			c.ServeJSON()
		} else {
			// Terjadi kesalahan lain
			c.Ctx.Output.SetStatus(http.StatusBadRequest)
			c.Data["json"] = map[string]string{"error": "Failed to retrieve user"}
			c.ServeJSON()
		}
		return
	}

	// Response user dalam format JSON
	c.Data["json"] = models.UserDefault(user)
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
	if err := utils.ValidateRequiredFields(&user, []string{"Name", "Email", "Password"}); err != nil {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]string{"error": err.Error()}
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
		c.Data["json"] = models.UserComplete(user)
	}
	c.ServeJSON()
}

func (c *UserController) Update() {
	// Ambil ID dari parameter URL
	id, errId := c.GetInt(":id")
	if errId != nil {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]string{"error": "Invalid user ID"}
		c.ServeJSON()
		return
	}

	var user models.User
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &user)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]string{"error": "Invalid JSON format"}
		c.ServeJSON()
		return
	}

	if err := utils.ValidateRequiredFields(&user, []string{"Name", "Email"}); err != nil {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}

	// Gunakan ORM untuk update data
	o := orm.NewOrm()
	newUserData := models.User{Id: id}

	// Cek apakah user ada
	if err := o.Read(&newUserData); err != nil {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]string{"error": "User not found"}
		c.ServeJSON()
		return
	}

	newUserData.Name = user.Name
	newUserData.Email = user.Email

	if _, err := o.Update(&newUserData); err != nil {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]string{"error": "Failed to update user"}
		c.ServeJSON()
		return
	}

	// Response sukses
	c.Data["json"] = models.UserDefault(newUserData)
	c.ServeJSON()
}

// Fungsi untuk menghapus user
func (c *UserController) Delete() {
	// Ambil ID dari parameter URL
	id, err := c.GetInt(":id")
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]string{"error": "Invalid user ID"}
		c.ServeJSON()
		return
	}

	// ORM instance
	o := orm.NewOrm()
	user := models.User{Id: id}

	// Periksa apakah user ada
	if err := o.Read(&user); err != nil {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]string{"error": "User not found"}
		c.ServeJSON()
		return
	}

	// Set DeletedAt dengan waktu saat ini
	now := time.Now()
	user.DeletedAt = &now

	// Update user dengan DeletedAt yang telah diubah
	if _, err := o.Update(&user, "DeletedAt"); err != nil {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]string{"error": "Failed to soft delete user"}
		c.ServeJSON()
		// c.CustomAbort(http.StatusInternalServerError, "Failed to soft delete user")
		return
	}

	// Response sukses
	c.Data["json"] = map[string]string{"message": "User deleted successfully"}
	c.ServeJSON()
}
