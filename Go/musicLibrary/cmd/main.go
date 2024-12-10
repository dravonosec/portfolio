package main

import (
	"fmt"
	"musicLibrary/internal/config"
	"musicLibrary/internal/database"
	"musicLibrary/internal/handler"
	"musicLibrary/internal/migrations"
	"musicLibrary/internal/repository"
	"musicLibrary/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {

	// start app
	// 1. load config
	cfg := config.LoadConfig()
	fmt.Println("Config loaded")

	// 2. init db
	dbCreator := database.NewDatabase(cfg)
	dbCreator.Create("music_db")

	// 3. run migrations;
	migrator := migrations.NewMigrator("file://../migarations", cfg.GetMusicDBConnString())
	migrator.RunMigration()

	// 4 init repo;
	repository := repository.NewRepository(cfg)
	fmt.Println("Repository loaded")

	// 5. init service
	service := service.NewMusicService(repository)
	fmt.Println("Service loaded")

	// 6. init handler
	handler := handler.NewHandler(service)
	fmt.Println("Handler loaded")

	// 7. init router
	router := gin.Default()
	handler.InitRoutes(router)

	router.Run(":8080")
	// start server
}
