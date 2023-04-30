package app

import "github.com/gin-gonic/gin"

func InitRouters(r *gin.Engine) {
	//
	r.GET("/ping", Ping)
	r.POST("/createUser", CreateUser)
	r.POST("/createTeam", CreateTeam)
	r.POST("/createEvent", CreateEvent)

	r.GET("/getUserByID", GetUserByID)
	r.GET("/getAllFormats", GetAllFormats)
	r.GET("/getAllRegions", GetAllRegions)
	r.GET("/getAllTags", GetAllTags)
	r.GET("/getAllGlobalTags", GetAllGlobalTags)
	r.GET("/getAllTagsByGlobalTag", GetAllTagsByGlobalTag)
	r.GET("/getAllUsers", GetAllUsers)
	r.GET("/getAllEvents", GetAllEvents)

	//r.GET("/getTeamsByUserID", GetTeamsByUserID)
	r.GET("/getTeamMembers", GetTeamMembers)

}
