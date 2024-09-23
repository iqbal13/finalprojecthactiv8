package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/iqbal13/finalprojecthactiv8/helpers"
	"github.com/iqbal13/finalprojecthactiv8/models"
	"github.com/iqbal13/finalprojecthactiv8/structs"
	"github.com/iqbal13/finalprojecthactiv8/usecase"
)

type userHandler struct {
	userUseCase usecase.UserUseCase
}

func NewUserHandler(routers *gin.Engine, userUseCase usecase.UserUseCase) {
	handler := &userHandler{userUseCase}

	router := routers.Group("/users")
	{
		router.POST("/register", handler.Register)
		router.POST("/login", handler.Login)
	}
}

func (handler *userHandler) Register(ctx *gin.Context) {
	var (
		user models.User
		err  error
	)

	if err = ctx.ShouldBindJSON(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "fail",
			Message: err.Error(),
		})

		return
	}

	if err = handler.userUseCase.Register(ctx.Request.Context(), &user); err != nil {
		if strings.Contains(err.Error(), "idx_users_username") {
			ctx.AbortWithStatusJSON(http.StatusConflict, helpers.ResponseMessage{
				Status:  "fail",
				Message: "Username Telah Digunakan",
			})

			return
		}

		if strings.Contains(err.Error(), "idx_users_email") {
			ctx.AbortWithStatusJSON(http.StatusConflict, helpers.ResponseMessage{
				Status:  "fail",
				Message: "Email Telah Digunakan",
			})

			return
		}

		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "fail",
			Message: err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusCreated, helpers.ResponseData{
		Status: "success",
		Data: structs.RegisteredUser{
			Age:      user.Age,
			Email:    user.Email,
			ID:       user.ID,
			Username: user.Username,
		},
	})
}
func (handler *userHandler) Login(ctx *gin.Context) {
	var (
		user  models.User
		err   error
		token string
	)

	if err = ctx.ShouldBindJSON(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "fail",
			Message: err.Error(),
		})

		return
	}

	if err = handler.userUseCase.Login(ctx.Request.Context(), &user); err != nil {
		if strings.Contains(err.Error(), "Kombinasi Email/Username dan Password Salah !") {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
				Status:  "Unauthenticated",
				Message: err.Error(),
			})

			return
		}

		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "Unauthenticated",
			Message: err.Error(),
		})

		return
	}

	if token = helpers.GenerateToken(user.ID, user.Email); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Unauthenticated",
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, helpers.ResponseData{
		Status: "success",
		Data: structs.LoggedinUser{
			Token: token,
		},
	})
}
