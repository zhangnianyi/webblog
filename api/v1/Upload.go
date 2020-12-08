package v1

import (
	"ginblog/modles"
	"ginblog/utils/errormessage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Upload(c *gin.Context){
	file,fileheader,_ :=c.Request.FormFile("file")
	fileSize :=fileheader.Size
	url,code :=modles.UploadFile(file,fileSize)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message":errormessage.GetErrorMessage(code),
		"url":url,
	})

}