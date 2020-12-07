package routes

import (
	"ginblog/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter()   {
	gin.SetMode(utils.AppMode)
	r :=gin.Default()
	route := r.Group("api/v1")
	{
		route.GET("/hello", func(c *gin.Context) {
			c.JSON(http.StatusOK,gin.H{
				"message":"你好啊 树先生",
			})

		})
	}

	r.Run()

}
