package main

import (
	"first-api/Config"
	"first-api/Models"
	"frist-api/Routes"
	"fmt"
)
var err error
func main() {
	  Config.DB, err = gorm.Open("mysql",Config.DbURL(Config.BuildDBConfig()))
	  if err!=nil {
		  fmt.Println("Status:", err)
	  }
	  defer Config.DB.Close()
	  Config.DB.AutoMigrate(&Models.User{})
	  r := Routes.SetupRouter()
	  r.Run()
}