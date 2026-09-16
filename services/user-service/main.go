package main

import (
	"log"
	"services/user-service/internal/config"
	"services/user-service/internal/user"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	_ = godotenv.Load()
}

func main() {
	db, err := config.InitDB()

	if err != nil {
		log.Fatalf("A database Error has occured %v", err)
	}

	repository := user.NewUserRepository(db)
	service := user.NewUserService(repository)
	handler := user.NewUserHandler(service)

	router := gin.Default()
	handler.Register(router)

	router.Run()
}
