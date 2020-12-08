package modles

import (
	"ginblog/utils/errormessage"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(100);not null" json:"username"`
	Password string  `gorm:"type:varchar(100);not null" json:"password"`
	Role int  `gorm:"type:int" json:"role"`  //0是管理员 1是阅读者
}
//查询用户是否存在
func Checkuser(name string)int{
	var users User
	DB.Select("id").Where("username=?",name).First(&users)
	if users.ID  >0{
		return errormessage.ERROR_USERNAME_USE
	}
	return errormessage.SUCCESS
}

//新增用户
func CreateUser(data *User)int{
	err := DB.Create(&data).Error
	if err != nil {
		return  errormessage.ERROR
	}
   return  errormessage.SUCCESS
}
func Getusers(pagesize int,pagenum int)[]*User{
	users :=make([]*User,0,100)
	//var users []User
	err =DB.Limit(pagesize).Offset((pagenum-1)*pagesize).Find(&users).Error
	if  err !=nil &&err ==gorm.ErrRecordNotFound{
		return  nil
	}
	return  users

}
func Edituser(id int){

}

