package app

import (
	"database/sql"
	"define-komandy/internal/structs"
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

// CreateUser returns id
func CreateUser(c *gin.Context) {

	// TODO add tags when creating user

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
		_, err := DB.Exec("INSERT INTO" + " users (name, nickname, rate, description, friends, logo, media, region_id, tags, mail, password) VALUES (" + str + ")")
		if err != nil {
			log.Print(err)
			c.JSON(500, gin.H{
				"server": -1,
			})
		} else {
			// User created
			var id int
			str := fmt.Sprintf("select user_id from users where mail='%s'", mail)
			row := DB.QueryRow(str)
			if err != nil {
				log.Print(err)
				_, _ = DB.Exec(fmt.Sprintf("delete from users where mail='%s'", mail))
				c.JSON(500, gin.H{
					"server": -1,
				})
			}
			err = row.Scan(&id)
			if err != nil {
				log.Print(err)
			}
			c.JSON(200, gin.H{
				"server": 1,
				"id":     id,
			})
		}
	}
}

func CreateTeam(c *gin.Context) {
	name := c.Query("name")
	description := c.Query("description")
	rules := c.Query("rules")
	place := c.Query("place")
	region := c.Query("region")

	// TODO: check for spam

	str := fmt.Sprintf("'%s', 0, '%s', '%s', '%s', '%s', %s, '%s', '%s', '%s'", name, description, rules, "", "", "current_timestamp", place, "{}", region)
	_, err := DB.Exec("INSERT" + " INTO team (name, rate, description, rules, logo, media, reg_date, place, tags, region_id) VALUES (" + str + ")")
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
	id := c.Query("id")
	row := DB.QueryRow(fmt.Sprintf("select * "+"from users "+"where user_id=%s", id))
	var user structs.User
	err := row.Scan(&user.UserId, &user.Name, &user.Nickname, &user.Rate, &user.Description, &user.Friends,
		&user.Logo, &user.Media, &user.Mail, &user.Password, &user.Tags, &user.Region)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"server": -1,
		})
	} else {
		c.JSON(200, gin.H{
			"user": user,
		})
	}
}

// CreateEvent returns inserted ID
func CreateEvent(c *gin.Context) {
	name := c.Query("name")
	description := c.Query("description")
	date := c.Query("date")
	format_id := c.Query("format_id")
	main_theme := c.Query("main_theme")
	place := c.Query("place")
	url := c.Query("url")
	tags := c.Query("tags")
	region_id := c.Query("region_id")
	creator_id := c.Query("creator_id")
	tagstr := "{" + tags + "}"

	// TAGS format (IDs): GET "1,6,9,3" = [1, 6, 9, 3]

	str := fmt.Sprintf("'%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', %s, %s, %s", name, description, date, main_theme, "", place,
		url, tagstr, format_id, region_id, creator_id)

	_, err := DB.Exec("INSERT " + "INTO " + "events (name, description, date, main_theme, media, place, url, tags, format_id, region_id, creator_id) VALUES (" + str + ")")
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"server": -1,
		})
	} else {
		var id int
		str := fmt.Sprintf("select event_id from events where creator_id='%s'", creator_id)
		row := DB.QueryRow(str)
		if err != nil {
			log.Print(err)
			_, _ = DB.Exec(fmt.Sprintf("delete from events where creator_id='%s'", creator_id))
			c.JSON(500, gin.H{
				"server": -1,
			})
		}
		err = row.Scan(&id)
		if err != nil {
			log.Print(err)
		}
		c.JSON(200, gin.H{
			"server": 1,
			"id":     id,
		})
	}
}

func GetAllFormats(c *gin.Context) {

}

func GetAllRegions(c *gin.Context) {

}

func GetAllTags(c *gin.Context) {

}
