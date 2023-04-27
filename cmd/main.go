package main

import (
	"database/sql"
	"define-komandy/internal/service"
	"log"
)

func main() {
	db := service.ConnectDB()
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	//testing-only:
	//service.DBFill(db)
}
