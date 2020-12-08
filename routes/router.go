package routes

import (
	v1 "ginblog/api/v1"
	"ginblog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter()   {
	gin.SetMode(utils.AppMode)
	r :=gin.Default()
	router := r.Group("api/v1")
	{
		//user模块的路由接口
		router.POST("/user/add",v1.Adduser)
		router.GET("/users",v1.Getusers)
		router.PUT("/user/:id",v1.Edituser)
		router.DELETE("/user/:id",v1.Deluser)

	}

	r.Run()

}
