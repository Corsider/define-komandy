package app

import "github.com/gin-gonic/gin"

func InitRouters(r *gin.Engine) {
	//
	r.GET("/ping", Ping)
	r.GET("/createUser", CreateUser)
	r.GET("/createTeam", CreateTeam)
	r.GET("/createEvent", CreateEvent)

	r.GET("/getUserByID", GetUserByID)
	r.GET("/getAllFormats", GetAllFormats)
	r.GET("/getAllRegions", GetAllRegions)
	r.GET("/getAllTags", GetAllTags)
}
