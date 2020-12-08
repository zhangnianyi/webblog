package modles

import (
	"fmt"
	"ginblog/utils/errormessage"
	"github.com/jinzhu/gorm"
)

type Artucle struct {
		gorm.Model
		Category Category  `gorm:"foreignkey:Cid"`
		Title string   `gorm:"type:varchar(100);not null" json:"title"`
		Cid int  `gorm:"type:int;not null" json:"cid"`
		Desr string  `gorm:"type:varchar(200)" json:"desr"`
		Content string `gorm:"type:longtext" json:"content"`
		Img string  `gorm:"type:varchar(100)" json:"img"`
}

//新增分类
func CreateArtucle(data *Artucle)int{
	err := DB.Create(&data).Error
	if err != nil {
		return  errormessage.ERROR
	}
	return  errormessage.SUCCESS
}
//查询单个文章信息
func GetArtinfobyid(id int)(Artucle,int){
	//var air =new(Artucle)
	var art  Artucle

	err :=DB.Debug().Preload("Category").Where("id = ?",id).First(&art).Error
	if  err !=nil {
		return art,errormessage.ERROR_ART_NOTEXIST
	}
	return  art,errormessage.SUCCESS
}
//id是分类的id
func GetArtbyCat(id int ,pagesize int,pagenum int)([]*Artucle,int){
	//var CatArtList []Artucle
	articleList :=make([]*Artucle,0,100)
	err =DB.Preload("Category").Limit(pagesize).Offset((pagenum-1)*pagesize).Where("cid = ?",id).Find(&articleList).Error
	if  err !=nil {
		return nil,errormessage.ERROR_ART_NOTEXIST
	}
	return articleList,errormessage.SUCCESS
}



func GetArtucle(pagesize int,pagenum int)([]*Artucle,int){
	articleList :=make([]*Artucle,0,100)
	//var users []User
	err =DB.Preload("Category").Limit(pagesize).Offset((pagenum-1)*pagesize).Find(&articleList).Error
	if  err !=nil &&err ==gorm.ErrRecordNotFound{
		return  nil,errormessage.ERROR
	}
	return  articleList,errormessage.SUCCESS

}
//查询单个文章 下次在做
func EditArtucle(id int,data *Artucle) int{
	art :=new(Artucle)
	var maps = make(map[string]interface{})
	maps["Title"] = data.Title
	maps["Cid"] = data.Cid
	maps["Desr"] = data.Desr
	maps["Content"] = data.Content
	maps["Img"] = data.Img
	err :=DB.Model(art).Where("id = ?",id).Update(maps).Error
	if err!=nil{
		fmt.Println(err)
		return errormessage.ERROR
	}
	return errormessage.SUCCESS

}

func DeleteArtucle(id int)int{
	var cate Artucle
	err :=DB.Where("id = ?",id).Delete(&cate).Error
	if err!=nil{
		fmt.Println(err)
		return errormessage.ERROR
	}
	return errormessage.SUCCESS
}
