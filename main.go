package main

import (
	"log"
	"simple-api-beego/database"
	_ "simple-api-beego/routers"

	_ "github.com/lib/pq"

	"github.com/beego/beego/orm"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	// database connection
	_, err := database.Connect()
	if err != nil {
		panic(err)
	}
	// Sinkronisasi database / migrasi, berdasarkan model yang diregistrasi
	err_syc := orm.RunSyncdb("default", false, true) // false: jangan hapus data lama, true: tampilkan log
	if err_syc != nil {
		log.Println("Error syncing database:", err_syc)
	}

	beego.Run()
}
