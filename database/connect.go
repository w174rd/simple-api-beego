package database

import (
	"fmt"
	"simple-api-beego/models"

	"github.com/beego/beego/orm"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/lib/pq"
)

func init() {
	// Register models
	orm.RegisterModel(new(models.User))

	orm.Debug = web.AppConfig.DefaultBool("isOrmDebug", false) // Aktifkan debug mode
}

func Connect() (db *orm.DB, err error) {
	// Register database
	orm.RegisterDriver("postgres", orm.DRPostgres)

	db_driver, _ := web.AppConfig.String("db.driver")
	db_user, _ := web.AppConfig.String("db.user")
	db_password, _ := web.AppConfig.String("db.password")
	db_name, _ := web.AppConfig.String("db.name")
	db_host, _ := web.AppConfig.String("db.host")
	db_port, _ := web.AppConfig.Int("db.port")
	db_sslmode, _ := web.AppConfig.String("db.sslmode")

	sqlconn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=%s", db_user, db_password, db_name, db_host, db_port, db_sslmode)
	orm.RegisterDataBase("default", db_driver, sqlconn)
	return
}
