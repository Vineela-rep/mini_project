package main

import (
	"Day4-5/Config"
	"Day4-5/Models"
	"Day4-5/Routes"
	"fmt"
	"github.com/jinzhu/gorm"
)
var err error
func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.User{})
	Config.DB.AutoMigrate(&Models.Product{})
	Config.DB.AutoMigrate(&Models.Order{})
	Config.DB.AutoMigrate(&Models.Retailer{})
	r := Routes.SetupRouter()
	//running
	r.Run()
}
