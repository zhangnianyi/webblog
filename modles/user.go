package modles

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(100);not null" json:"username"`
	Password string  `gorm:"type:varchar(100);not null" json:"password"`
	Role int  `gorm:"type:int" json:"role"`  //0是管理员 1是阅读者
}