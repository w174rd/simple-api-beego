package models

import (
	"time"
)

type User struct {
	Id        int        `json:"id" orm:"auto"`
	Name      string     `json:"name" orm:"size(100)"`
	Email     string     `json:"email" orm:"size(100)"`
	DeletedAt *time.Time `json:"deleted_at" orm:"null;type(timestamp)"`
	CreatedAt time.Time  `json:"created_at" orm:"auto_now_add;type(timestamp)"`
	UpdateAt  time.Time  `json:"update_at" orm:"auto_now;type(timestamp)"`
}
