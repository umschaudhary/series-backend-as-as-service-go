package main

import (
	"blog/api/controllers"
	"blog/api/repositories"
	"blog/api/routes"
	"blog/api/services"
	"blog/infrastructure"
	"blog/models"
)

func main() {

	router := infrastructure.NewGinRouter() //router has been initialized and configured
	db := infrastructure.NewDatabase() // databse has been initialized and configured
	postRepository := repositories.NewPostRepository(db) // repository are being setup
	postService := services.NewPostService(postRepository) // service are being setup
	postController := controllers.NewPostController(postService) // controller are being set up
	postRoute := routes.NewPostRoute(postController, router) // post routes are initialized
	postRoute.Setup() // post routes are being setup

	db.DB.AutoMigrate(&models.Post{}) // migrating Post model to datbase table
	router.Gin.Run(":8000") //server started on 8000 port
}
