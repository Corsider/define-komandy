package app

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

var DB *sql.DB

func CreateUser(c *gin.Context) {
	name := c.Query("name")
	nickname := c.Query("nickname")
	password := c.Query("password")
	description := c.Query("description")
	region := c.Query("region")
	mail := c.Query("mail")

	str := fmt.Sprintf("\\'%s\\' \\'%s\\' \\'%d\\' \\'%s\\' \\'%s\\' \\'%s\\' \\'%s\\' \\'%s\\' \\'%s\\' \\'%s\\' \\'%s\\'",
		name, nickname, 0, description, "{}", "", "", region, "{}", mail, HashPassword(password))
	_, err := DB.Exec("INSERT INTO users (name, nickname, rate, description, friends, logo, media, region, tags, mail, password) VALUES (" + str + ")")
	if err != nil {
		log.Print(err)
	}
}
