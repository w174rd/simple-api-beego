package models

import (
	"time"
)

type User struct {
	Id        int        `orm:"auto"`
	Name      string     `orm:"size(100)"`
	Email     string     `orm:"size(100)"`
	DeletedAt *time.Time `orm:"null;type(timestamp)"`
	CreatedAt time.Time  `orm:"auto_now_add;type(timestamp)"`
	UpdateAt  time.Time  `orm:"auto_now_add;type(timestamp)"`
}
