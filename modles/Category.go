package modles

import (
	"fmt"
	"ginblog/utils/errormessage"
	"github.com/jinzhu/gorm"
)

type Category struct {
	ID uint `gorm:"primary_key;auto_increment"  json:"id"`
	Name string `gorm:"type:varchar(100);not null"  json:"name"`
}

func CheckCategory(name string)int{
	var cate Category
	DB.Select("id").Where("name=?",name).First(&cate)
	if cate.ID  >0{
		return errormessage.ERROR_CATRNAME_USERD
	}
	return errormessage.SUCCESS
}

//新增分类
func CreateCategory(data *Category)int{
	err := DB.Create(&data).Error
	if err != nil {
		return  errormessage.ERROR
	}
	return  errormessage.SUCCESS
}
func GetCategory(pagesize int,pagenum int)[]*Category{
	cate :=make([]*Category,0,100)
	//var users []User
	err =DB.Limit(pagesize).Offset((pagenum-1)*pagesize).Find(&cate).Error
	if  err !=nil &&err ==gorm.ErrRecordNotFound{
		return  nil
	}
	return  cate

}

func EditCategory(id int,data *Category) int{
	cate :=new(Category)
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	err :=DB.Model(cate).Where("id = ?",id).Update(maps).Error
	if err!=nil{
		fmt.Println(err)
		return errormessage.ERROR
	}
	return errormessage.SUCCESS

}

func DeleteCategory(id int)int{
	var cate Category
	err :=DB.Where("id = ?",id).Delete(&cate).Error
	if err!=nil{
		fmt.Println(err)
		return errormessage.ERROR
	}
	return errormessage.SUCCESS
}

