package config

import "github.com/jinzhu/gorm"

var (
	db *gorm.DB
)

func Connect() {
	// Add mysql url
	d, err := gorm.Open("mysql", "rakesh:rakesh@12@/simplego?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
