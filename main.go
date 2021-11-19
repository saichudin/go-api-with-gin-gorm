package main

import (
	"api-crud-with-gin/handler"
	"api-crud-with-gin/user"

	"log"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:8888)/go_crud?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB connection error")
	}

	db.AutoMigrate(&user.User{})

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	bookHandler := handler.NewBookHandler(userService)

	// users, err := userRepository.FindAll()

	// for _, user := range users {
	// 	fmt.Println("username :", user.Username)
	// }

	route := gin.Default()
	route.GET("/user", bookHandler.RootHandler)
	route.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
