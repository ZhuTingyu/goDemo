package main

import (
	"demo/database"
	"demo/user"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Person struct {
	ID        uint   `json:"id"  gorm:"id"`
	FirstName string `json:"firstname" gorm:"first_name"`
	LastName  string `json:"lastname" gorm:"last_name"`
	City      string `json:"city" gorm:"city"`
}

func main() {
	database.Init()
	r := gin.Default()
	r.GET("/user", user.GetById)
	r.Run(":8800")
}
