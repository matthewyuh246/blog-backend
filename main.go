package main

import (
	"github.com/matthewyuh246/blogbackend/controller"
	"github.com/matthewyuh246/blogbackend/db"
	repository "github.com/matthewyuh246/blogbackend/repositroy"
	"github.com/matthewyuh246/blogbackend/router"
	"github.com/matthewyuh246/blogbackend/usecase"
	"github.com/matthewyuh246/blogbackend/validator"
)

func main() {
	db := db.NewDB()
	userValidator := validator.NewUserValidator()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	userController := controller.NewUserController(userUsecase)

	blogRepository := repository.NewBlogRepository(db)
	blogUsecase := usecase.NewBlogUsecase(blogRepository)
	blogController := controller.NewBlogController(blogUsecase)

	imageRepository := repository.NewImageRepository("http://localhost:8080/api/blog/uploads/")
	imageUsecase := usecase.NewImageUsecase(imageRepository)
	imageController := controller.NewImageController(imageUsecase)

	e := router.NewRouter(userController, blogController, imageController)
	e.Logger.Fatal(e.Start(":8080"))
}
