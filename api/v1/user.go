package v1

import (
	"ginblog/modles"
	"ginblog/utils/errormessage"
	"ginblog/utils/validator"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)
var code int

//添加用
func Adduser(c *gin.Context){
	data  :=new(modles.User)
	var msg string
	_ =c.ShouldBindJSON(data)
	msg,code =validator.Validatroy(data)
	if code !=errormessage.SUCCESS{
		c.JSON(http.StatusOK,gin.H{
			"status":code,
			"message":msg,
		})
		return
	}
	code =modles.Checkuser(data.Username)
	if code == errormessage.SUCCESS{
		modles.CreateUser(data)
	}
	if code == errormessage.ERROR_USERNAME_USE{
		code =errormessage.ERROR_USERNAME_USE

	}
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data":data,
		"message":errormessage.GetErrorMessage(code),
	})

}
//查询用户：查询单个用户
//
//查询用户列表
func Getusers(c *gin.Context){
	pagesize,_ :=strconv.Atoi(c.Query("pagesize"))
	pagenum,_ :=strconv.Atoi(c.Query("pagenum"))
	if pagesize ==0{
		pagesize = -1
	}
	if pagenum ==0{
		pagenum =-1
	}
	data,total :=modles.Getusers(pagesize,pagenum)
	code = errormessage.SUCCESS
	c.JSON(http.StatusOK,gin.H{
		"STATUS":code,
		"data":data,
		"message":errormessage.GetErrorMessage(code),
		"total":total,
	})

}

//删除用户
func Deluser(c *gin.Context){
	id,_ :=strconv.Atoi(c.Param("id"))
	code =modles.Deleteuser(id)
	c.JSON(http.StatusOK,gin.H{
		"STATUS":code,
		"message":errormessage.GetErrorMessage(code),
	})
}
//编辑用户
func Edituser(c *gin.Context){

	var data modles.User
	id,_ :=strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data )
	code =modles.Checkuser(data.Username)
	if code == errormessage.SUCCESS{
		modles.Edituser(id,&data)
	}
	if code == errormessage.ERROR_USERNAME_USE{
		c.Abort()
	}
	c.JSON(http.StatusOK,gin.H{
		"STATUS":code,
		"message":errormessage.GetErrorMessage(code),
	})


}
