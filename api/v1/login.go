package v1

import (
	"ginblog/middleware"
	"ginblog/modles"
	"ginblog/utils/errormessage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context){
	 data :=new(modles.User)
	 c.ShouldBindJSON(data)
	 code=modles.CheckLogin(data.Username,data.Password)
	 var token string
	 if code == errormessage.SUCCESS{
	 	token,_ =middleware.GenToken(data.Username)
	 }
	 c.JSON(http.StatusOK,gin.H{
	 	"status":code,
	 	"message":errormessage.GetErrorMessage(code),
	 	"token":token,

	 })

}