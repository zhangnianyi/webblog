package v1

import (
	"fmt"
	"ginblog/modles"
	"ginblog/utils/errormessage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AddCategory(c *gin.Context){
	cate  :=new(modles.Category)
	_ =c.ShouldBindJSON(cate)
	code =modles.CheckCategory(cate.Name)
	if code == errormessage.SUCCESS{
		modles.CreateCategory(cate)
	}
	if code == errormessage.ERROR_CATRNAME_USERD{
		code =errormessage.ERROR_CATRNAME_USERD

	}
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data":cate,
		"message":errormessage.GetErrorMessage(code),
	})

}
//查询用户：查询单个用户
//
//查询用户列表
func GetCategory(c *gin.Context){
	pagesize,_ :=strconv.Atoi(c.Query("pagesize"))
	pagenum,_ :=strconv.Atoi(c.Query("pagenum"))
	if pagesize ==0{
		pagesize = -1
	}
	if pagenum ==0{
		pagenum =-1
	}
	data ,total:=modles.GetCategory(pagesize,pagenum)
	code = errormessage.SUCCESS
	c.JSON(http.StatusOK,gin.H{
		"STATUS":code,
		"data":data,
		"message":errormessage.GetErrorMessage(code),
		"total":total,
	})

}

//删除用户
func DelCategory(c *gin.Context){
	id,_ :=strconv.Atoi(c.Param("id"))
	code =modles.DeleteCategory(id)
	c.JSON(http.StatusOK,gin.H{
		"STATUS":code,
		"message":errormessage.GetErrorMessage(code),
	})
}
//编辑用户
func EditCategory(c *gin.Context){

	var data modles.Category
	id,_ :=strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)

	fmt.Println("去进行存在校验")
	code =modles.CheckCategory(data.Name)
	fmt.Println(data.ID)
	fmt.Println(data.Name)
	if code == errormessage.SUCCESS{
		fmt.Println("数据库校验完成")
		modles.EditCategory(id,&data)
	}
	if code == errormessage.ERROR_CATRNAME_USERD{
		c.Abort()
	}
	c.JSON(http.StatusOK,gin.H{
		"STATUS":code,
		"message":errormessage.GetErrorMessage(code),
	})


}
