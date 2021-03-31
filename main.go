package main

import (
	"dasargo/Config"
	"dasargo/Models"
	"dasargo/Routes"
	"github.com/jinzhu/gorm"
	"fmt"
)

var err error

func main() {
	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildConfig()))
	if err != nil {
		fmt.Println("Status :", err)
	}
	defer config.DB.Close()
	config.DB.AutoMigrate(&models.User{})
	
	router := routes.SetUpRouter()
	router.Run()
}
