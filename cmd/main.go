package main

import (
	"database/sql"
	"define-komandy/internal/app"
	"define-komandy/internal/service"
	"define-komandy/internal/structs"
	"github.com/gin-gonic/gin"
	"log"
)

var DB *sql.DB

func main() {
	DB = service.ConnectDB()
	err := DB.Ping()
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(DB)
	app.DB = DB
	r := gin.Default()

	err = r.Run(service.First(service.ReadYaml[structs.Config]("/config/config.yaml")).Host)
	if err != nil {
		log.Fatal(err)
	}
}
