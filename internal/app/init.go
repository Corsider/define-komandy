package app

import "github.com/gin-gonic/gin"

func InitRouters(r *gin.Engine) {
	//
	r.GET("/ver", CreateUser)
}
