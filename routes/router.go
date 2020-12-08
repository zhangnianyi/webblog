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
		//分类模块的路由接口
		router.POST("/category/add",v1.AddCategory)
		router.GET("/categorys",v1.GetCategory)
		router.PUT("/category/:id",v1.EditCategory)
		router.DELETE("/category/:id",v1.DelCategory)
		//wenzhang
		router.POST("/artucle/add",v1.AddArtucle)
		router.GET("/artucles",v1.GetArtucle)
		router.GET("/artucles/conlist/:id",v1.GetcateArt)
		router.GET("/artucles/list/:id",v1.Getaircinfobyid)
		router.PUT("/artucle/:id",v1.EditArtucle)
		router.DELETE("/artucle/:id",v1.DelArtucle)

	}

	r.Run()

}
