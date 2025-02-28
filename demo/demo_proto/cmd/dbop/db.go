package main

import (
	"fmt"

	"github.com/cloudwego/biz-demo/gomall/demo/demo_proto/biz/dal"
	"github.com/cloudwego/biz-demo/gomall/demo/demo_proto/biz/dal/mysql"
	"github.com/cloudwego/biz-demo/gomall/demo/demo_proto/biz/model"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	dal.Init()
	// CURD

	// Create
	//mysql.DB.Create(&model.User{Email: "demo@example.com", Password: "jfiojffjsoij"})

	// Update
	//mysql.DB.Model(&model.User{}).Where("email = ?", "demo@example.com").Update("password", "22222222")

	// Read
	var row model.User
	mysql.DB.Model(&model.User{}).Where("email = ?", "demo@example.com").First(&row)
	
	fmt.Printf("row: %+v\n", row)

	// Delete
	// mysql.DB.Where("email = ?", "demo@example.com").Delete(&model.User{})

	// mysql.DB.Unscoped().Where("email = ?", "demo@example.com").Delete(&model.User{})
}