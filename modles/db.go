package modles

import (
	"fmt"
	"ginblog/utils"
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var (
	DB  *gorm.DB
	err error
)
func InitDb(){
	DB ,err =gorm.Open(utils.Db,fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.Dbuser,
		utils.Dbpwd,
		utils.DbHost,
		utils.DbPort,
		utils.Dbname,
		))
	if err !=nil{
		fmt.Println("connect mysql faild please checking",err)
		return
	}
	DB.AutoMigrate(&User{},&Artucle{},&Category{})
	//DB.SingularTable(true)
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
	DB.DB().SetConnMaxLifetime(10*time.Second)
}