package routes

import (
	v1 "ginblog/api/v1"
	"ginblog/middleware"
	"ginblog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter()   {
	gin.SetMode(utils.AppMode)
	r :=gin.Default()
	auth := r.Group("api/v1")
	auth.Use(middleware.Jwttoken())
	{
		//user模块的路由接口


		auth.PUT("/user/:id",v1.Edituser)
		auth.DELETE("/user/:id",v1.Deluser)
		//分类模块的路由接口
		auth.POST("/category/add",v1.AddCategory)

		auth.PUT("/category/:id",v1.EditCategory)
		auth.DELETE("/category/:id",v1.DelCategory)
		//wenzhang
		auth.POST("/artucle/add",v1.AddArtucle)

		auth.PUT("/artucle/:id",v1.EditArtucle)
		auth.DELETE("/artucle/:id",v1.DelArtucle)
		//上传文件
		auth.POST("/upload",v1.Upload)

	}
	router := r.Group("api/v1")
	{
		router.POST("/user/add",v1.Adduser)
		router.GET("/users",v1.Getusers)
		router.GET("/categorys",v1.GetCategory)
		router.GET("/artucles",v1.GetArtucle)
		router.GET("/artucles/conlist/:id",v1.GetcateArt)
		router.GET("/artucles/list/:id",v1.Getaircinfobyid)
		router.POST("/login",v1.Login)
	}


	r.Run()

}
