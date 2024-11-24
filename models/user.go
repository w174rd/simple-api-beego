package models

import (
	"time"
)

type User struct {
	Id        int        `json:"id" orm:"auto"`
	Name      string     `json:"name" orm:"size(100)"`
	Email     string     `json:"email" orm:"size(100)"`
	Password  string     `json:"password,omitempty" orm:"null;size(100)"`
	Token     string     `json:"token,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" orm:"null;type(timestamp)"`
	CreatedAt *time.Time `json:"created_at,omitempty" orm:"auto_now_add;type(timestamp)"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" orm:"auto_now;type(timestamp)"`
}

func UserDefault(user User) User {
	return User{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
		// UpdatedAt: user.UpdatedAt,
		// CreatedAt: user.CreatedAt,
		// DeletedAt: user.DeletedAt,
	}
}

func UserComplete(user User) User {
	return User{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		UpdatedAt: user.UpdatedAt,
		CreatedAt: user.CreatedAt,
		DeletedAt: user.DeletedAt,
	}
}
