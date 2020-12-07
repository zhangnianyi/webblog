package routes

import (
	"ginblog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter()   {
	gin.SetMode(utils.AppMode)
	r :=gin.Default()
	v1 := r.Group("api/v1")
	{
		//user模块的路由接口
	}

	r.Run()

}
