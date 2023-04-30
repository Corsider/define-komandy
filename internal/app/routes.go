package app

import (
	"database/sql"
	"define-komandy/internal/structs"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
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

	var usr structs.User
	_ = c.BindJSON(&usr)

	//name := usr.Name
	//nickname := c.Query("nickname")
	//password := c.Query("password")
	//description := c.Query("description")
	//region := c.Query("region")
	//mail := c.Query("mail")

	query := fmt.Sprintf("SELECT nickname FROM users WHERE nickname='%s' OR mail='%s'", usr.Nickname, usr.Mail)
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
			usr.Name, usr.Nickname, 0, usr.Description, "{}", "", "", usr.RegionID, "{}", usr.Mail, HashPassword(usr.Password))
		_, err := DB.Exec("INSERT INTO" + " users (name, nickname, rate, description, friends, logo, media, region_id, tags, mail, password) VALUES (" + str + ")")
		if err != nil {
			log.Print(err)
			c.JSON(500, gin.H{
				"server": -1,
			})
		} else {
			// User created
			var id int
			str := fmt.Sprintf("select user_id from users where mail='%s'", usr.Mail)
			row := DB.QueryRow(str)
			if err != nil {
				log.Print(err)
				_, _ = DB.Exec(fmt.Sprintf("delete from users where mail='%s'", usr.Mail))
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
	//name := c.Query("name")
	//description := c.Query("description")
	//rules := c.Query("rules")
	//place := c.Query("place")
	//region := c.Query("region")

	var team structs.Team
	_ = c.BindJSON(&team)

	// TODO: check for spam

	str := fmt.Sprintf("'%s', 0, '%s', '%s', '%s', '%s', %s, '%s', '%s', '%s'", team.Name, team.Description, team.Rules, "", "", "current_timestamp", team.Place, "{}", team.RegionID)
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
		&user.Logo, &user.Media, &user.Mail, &user.Password, &user.Tags, &user.RegionID)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"server": -1,
		})
	} else {
		user.Password = ""
		c.JSON(200, gin.H{
			"user": user,
		})
	}
}

// CreateEvent returns inserted ID
func CreateEvent(c *gin.Context) {
	//name := c.Query("name")
	//description := c.Query("description")
	//date := c.Query("date")
	//format_id := c.Query("format_id")
	//main_theme := c.Query("main_theme")
	//place := c.Query("place")
	//url := c.Query("url")
	//tags := c.Query("tags")
	//region_id := c.Query("region_id")
	//creator_id := c.Query("creator_id")
	//tagstr := "{" + tags + "}"
	var event structs.Event
	_ = c.BindJSON(&event)
	tagstr := "{"
	for i, el := range event.Tags {
		if i != len(event.Tags)-1 {
			tagstr += strconv.Itoa(int(el)) + ","
		} else {
			tagstr += strconv.Itoa(int(el)) + "}"
		}
	}
	//fmt.Println(tagstr)

	// TODO: ADD TO TEAM_EVENTS
	// TAGS format (IDs): GET "1,6,9,3" = [1, 6, 9, 3]

	str := fmt.Sprintf("'%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', %d, %d, %d", event.Name, event.Description, event.Date, event.MainTheme, "", event.Place,
		event.Url, tagstr, event.FormatID, event.RegionID, event.CreatorID)
	var eventID int
	err := DB.QueryRow("INSERT " + "INTO " + "events (name, description, date, main_theme, media, place, url, tags, format_id, region_id, creator_id) VALUES (" + str + ") RETURNING event_id").Scan(&eventID)
	//fmt.Println(res.LastInsertId())
	//fmt.Println(res)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"server": -1,
		})
	} else {
		/*
			var id int
			str := fmt.Sprintf("select event_id from events where creator_id='%d'", event.CreatorID)
			row := DB.QueryRow(str)
			if err != nil {
				log.Print(err)
				_, _ = DB.Exec(fmt.Sprintf("delete from events where creator_id='%d'", event.CreatorID))
				c.JSON(500, gin.H{
					"server": -1,
				})
			}
			err = row.Scan(&id)
			if err != nil {
				log.Print(err)
			}
		*/
		c.JSON(200, gin.H{
			"server": 1,
			"id":     eventID,
		})
	}
}

func GetAllFormats(c *gin.Context) {
	formats := []structs.Format{}
	rows, err := DB.Query("SELECT * FROM format")
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"server": -1,
		})
	} else {
		for rows.Next() {
			var form structs.Format
			_ = rows.Scan(&form.FormatId, &form.Format)
			formats = append(formats, form)
		}
		c.JSON(200, gin.H{
			"formats": formats,
		})
	}
}

func GetAllRegions(c *gin.Context) {
	regions := []structs.Region{}
	rows, err := DB.Query("SELECT * FROM region")
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"server": -1,
		})
	} else {
		for rows.Next() {
			var reg structs.Region
			_ = rows.Scan(&reg.RegionID, &reg.CountryID, &reg.RegionName)
			regions = append(regions, reg)
		}
		c.JSON(200, gin.H{
			"regions": regions,
		})
	}
}

func GetAllTags(c *gin.Context) {
	tags := []structs.Tag{}
	rows, err := DB.Query("SELECT * FROM tag")
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"server": -1,
		})
	} else {
		for rows.Next() {
			var tagg structs.Tag
			_ = rows.Scan(&tagg.TagID, &tagg.Activity, &tagg.GlobalTagID)
			tags = append(tags, tagg)
		}
		c.JSON(200, gin.H{
			"tag": tags,
		})
	}
}

func GetAllGlobalTags(c *gin.Context) {

}

func GetAllTagsByGlobalTag(c *gin.Context) {

}

func GetAllUsers(c *gin.Context) {
	users := []structs.User{}
	rows, err := DB.Query("SELECT * FROM users")
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"server": -1,
		})
	} else {
		for rows.Next() {
			var user structs.User
			_ = rows.Scan(&user.UserId, &user.Name, &user.Nickname, &user.Rate, &user.Description, &user.Friends,
				&user.Logo, &user.Media, &user.Mail, &user.Password, &user.Tags, &user.RegionID)
			user.Mail = ""
			user.Password = ""
			user.Name = ""
			users = append(users, user)
		}
		c.JSON(200, gin.H{
			"users": users,
		})
	}
}

func GetAllEvents(c *gin.Context) {
	events := []structs.Event{}
	rows, err := DB.Query("SELECT * FROM events")
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"server": -1,
		})
	} else {
		for rows.Next() {
			var event structs.Event
			_ = rows.Scan(&event.EventID, &event.Name, &event.Description, &event.Date, &event.MainTheme,
				&event.Media, &event.Place, &event.Url, &event.Tags, &event.FormatID, &event.RegionID,
				&event.CreatorID)
			events = append(events, event)
		}
		c.JSON(200, gin.H{
			"events": events,
		})
	}
}

//func GetTeamsByUserID(c *gin.Context) {
//	id := c.Query("id")
//	rows, err := DB.Query("SELECT * FROM user_team")
//	if err != nil {
//		log.Println(err)
//		c.JSON(500, gin.H{
//			"server": -1,
//		})
//	} else {
//		for rows.Next() {
//			var
//		}
//	}
//}

func GetTeamMembers(c *gin.Context) {

}
