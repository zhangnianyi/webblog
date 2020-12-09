package modles

import (
	"encoding/base64"
	"fmt"
	"ginblog/utils/errormessage"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/scrypt"
	"log"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(100);not null" json:"username" validata:"required,min=4,max=12"`
	Password string  `gorm:"type:varchar(100);not null" json:"password" validata:"required,min=4,max=20`
	Role int  `gorm:"type:int;DEFAULT:2" json:"role" validata:"required,get=2`  //0是管理员 1是阅读者
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
	data.Password=ScryptPw(data.Password)
	err := DB.Create(&data).Error
	if err != nil {
		return  errormessage.ERROR
	}
   return  errormessage.SUCCESS
}
func Getusers(pagesize int,pagenum int)([]*User,int){
	users :=make([]*User,0,100)
	//var users []User
	var total int
	err =DB.Limit(pagesize).Offset((pagenum-1)*pagesize).Find(&users).Count(&total).Error
	if  err !=nil &&err ==gorm.ErrRecordNotFound{
		return  nil,0
	}
	return  users,total

}

func Deleteuser(id int)int{
	var users User
	err :=DB.Where("id = ?",id).Delete(&users).Error
	if err!=nil{
		fmt.Println(err)
		return errormessage.ERROR
	}
	return errormessage.SUCCESS
}

func Edituser(id int,data *User) int{
	user :=new(User)
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] =data.Role
	err :=DB.Model(user).Where("id = ?",id).Update(maps).Error
	if err!=nil{
		fmt.Println(err)
		return errormessage.ERROR
	}
	return errormessage.SUCCESS

}
//密码加密
func ScryptPw(password string)string{
	const   (
		Keylen =8
	)
	salt :=make([]byte,8)
	salt = []byte{4,56,32,65,132,131,22,21}
	Hashpwd,err :=scrypt.Key([]byte(password),salt,16384,8,1,Keylen)
	if err !=nil{
		log.Fatal(errormessage.ERROR)
	}
	FPW :=base64.StdEncoding.EncodeToString(Hashpwd)
	return FPW
}


//登录
func CheckLogin(username string,password string)int{
	user :=new(User)
	DB.Where("username =?",username).First(user)
	if user.ID ==0{
		 return errormessage.ERROR_USER_NOTEXIST
	}
	if ScryptPw(password) !=user.Password{
	return  errormessage.ERROR_PASSWORD_WRONG
	}
	if user.Role !=1{
		return  errormessage.ERROR_USER_NO_RIGHT
	}
	return  errormessage.SUCCESS


}
