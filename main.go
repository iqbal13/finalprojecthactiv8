package main

import (
	"log"
	"os"
	"time"

	"github.com/iqbal13/finalprojecthactiv8/config"
	"github.com/iqbal13/finalprojecthactiv8/controllers"
	"github.com/iqbal13/finalprojecthactiv8/repository"
	"github.com/iqbal13/finalprojecthactiv8/usecase"
	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

func main() {

	os.Setenv("TZ", "Asia/Jakarta")
	time.LoadLocation("Asia/Jakarta")

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Failed Load Environment")
	}

	config.ConnectDB()
	routers := gin.Default()

	db := config.DB

	routers.Use(func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Content-Type", "application/json")
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Max-Age", "86400")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, UPDATE")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(200)
		} else {
			ctx.Next()
		}
	})

	userRepository := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepository)
	controllers.NewUserHandler(routers, userUseCase)

	photoRepository := repository.NewPhotoRepository(db)
	photoUseCase := usecase.NewPhotoUseCase(photoRepository)
	controllers.NewPhotoHandler(routers, photoUseCase)

	commentRepository := repository.NewCommentRepository(db)
	commentUseCase := usecase.NewCommentUseCase(commentRepository)
	controllers.NewCommentHandler(routers, commentUseCase, photoUseCase)

	socialMediaRepository := repository.NewSocialMediaRepository(db)
	socialMediaUseCase := usecase.NewSocialMediaUseCase(socialMediaRepository)
	controllers.NewSocialMediaHandler(routers, socialMediaUseCase)

	routers.Run(os.Getenv("APP_PORT"))

}
