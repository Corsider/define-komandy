package service

import (
	"database/sql"
	"define-komandy/internal/structs"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func ConnectDB() *sql.DB {
	data, err := ReadYaml[structs.Config]("/config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	secret, err := ReadYaml[structs.SecretConfig]("/config/secret.yaml")
	if err != nil {
		log.Fatal(err)
	}
	connStr := fmt.Sprintf("user= %s password= %s dbname= %s sslmode=disable", secret.User, secret.Password, data.DbName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
