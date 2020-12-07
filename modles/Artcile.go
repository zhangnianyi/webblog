package modles

import "github.com/jinzhu/gorm"

type Artucle struct {
		gorm.Model
		Category Category
		Title string   `gorm:"type:varchar(100);not null" json:"title"`
		Cid int  `gorm:"type:int;not null" json:"cid"`
		Desr string  `gorm:"type:varchar(200)" json:"desr"`
		Content string `gorm:"type:longtext" json:"content"`
		Img string  `gorm:"type:varchar(100)" json:"img"`
}