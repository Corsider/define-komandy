package app

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

var DB *sql.DB

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"server": "1",
	})
}

func CreateUser(c *gin.Context) {
	name := c.Query("name")
	nickname := c.Query("nickname")
	password := c.Query("password")
	description := c.Query("description")
	region := c.Query("region")
	mail := c.Query("mail")

	query := fmt.Sprintf("SELECT nickname FROM users WHERE nickname='%s' OR mail='%s'", nickname, mail)
	rows, err := DB.Query(query)
	if err != nil {
		log.Println(err)
	}

	counter := 0
	for rows.Next() {
		counter += 1
	}

	if counter != 0 {
		c.JSON(403, gin.H{
			"server": 0,
		})
	} else {
		str := fmt.Sprintf("'%s', '%s', %d, '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s'",
			name, nickname, 0, description, "{}", "", "", region, "{}", mail, HashPassword(password))
		_, err := DB.Exec("INSERT INTO" + " users (name, nickname, rate, description, friends, logo, media, region, tags, mail, password) VALUES (" + str + ")")
		if err != nil {
			log.Print(err)
			c.JSON(500, gin.H{
				"server": -1,
			})
		} else {
			c.JSON(200, gin.H{
				"server": 1,
			})
		}
	}
}

func CreateTeam(c *gin.Context) {
	name := c.Query("name")
	description := c.Query("description")
	rules := c.Query("rules")
	place := c.Query("place")

	// TODO: check for spam

	str := fmt.Sprintf("'%s', 0, '%s', '%s', '%s', '%s', %s, '%s', '%s'", name, description, rules, "", "", "current_timestamp", place, "{}")
	_, err := DB.Exec("INSERT" + " INTO team (name, rate, description, rules, logo, media, reg_date, place, tags) VALUES (" + str + ")")
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"server": -1,
		})
	} else {
		c.JSON(200, gin.H{
			"server": 1,
		})
	}
}

func GetUserByID(c *gin.Context) {
	//id := c.Query("id")
	//str := fmt.Sprintf()
}
