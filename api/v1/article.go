package v1

import (
	"ginblog/modles"
	"ginblog/utils/errormessage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AddArtucle(c *gin.Context){
	data  :=new(modles.Artucle)
	_ =c.ShouldBindJSON(data)
	code =modles.CreateArtucle(data)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data":data,
		"message":errormessage.GetErrorMessage(code),
	})

}
//查询用户：查询单个用户
//
//查询用户列表
func GetArtucle(c *gin.Context){
	pagesize,_ :=strconv.Atoi(c.Query("pagesize"))
	pagenum,_ :=strconv.Atoi(c.Query("pagenum"))
	if pagesize ==0{
		pagesize = -1
	}
	if pagenum ==0{
		pagenum =-1
	}
	data,code,toal:=modles.GetArtucle(pagesize,pagenum)
	c.JSON(http.StatusOK,gin.H{
		"STATUS":code,
		"data":data,
		"message":errormessage.GetErrorMessage(code),
		"total":toal,
	})
}

//删除用户
func DelArtucle(c *gin.Context){
	id,_ :=strconv.Atoi(c.Param("id"))
	code =modles.DeleteArtucle(id)
	c.JSON(http.StatusOK,gin.H{
		"STATUS":code,
		"message":errormessage.GetErrorMessage(code),
	})
}
//编辑用户
func EditArtucle(c *gin.Context){

	var data modles.Artucle
	id,_ :=strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)

	code=modles.EditArtucle(id,&data)
	c.JSON(http.StatusOK,gin.H{
		"STATUS":code,
		"message":errormessage.GetErrorMessage(code),
	})


}
//查询单个文章
func  Getaircinfobyid(c *gin.Context)  {
	id,_ :=strconv.Atoi(c.Param("id"))
	data,code :=modles.GetArtinfobyid(id)
	c.JSON(http.StatusOK,gin.H{
		"STATUS":code,
		"data":data,
		"message":errormessage.GetErrorMessage(code),
	})
	}

//查询分类下面的所有文章
func GetcateArt(c *gin.Context){
	pagesize,_ :=strconv.Atoi(c.Query("pagesize"))
	pagenum,_ :=strconv.Atoi(c.Query("pagenum"))
	if pagesize ==0{
		pagesize = -1
	}
	if pagenum ==0{
		pagenum =-1
	}
	id,_ :=strconv.Atoi(c.Param("id"))
	data,code,total :=modles.GetArtbyCat(id,pagesize,pagesize)
	c.JSON(http.StatusOK,gin.H{
		"STATUS":code,
		"data":data,
		"message":errormessage.GetErrorMessage(code),
		"total":total,
	})
}