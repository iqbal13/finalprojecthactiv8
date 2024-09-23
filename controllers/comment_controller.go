package controllers

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/iqbal13/finalprojecthactiv8/helpers"
	middleware "github.com/iqbal13/finalprojecthactiv8/middlewares"
	"github.com/iqbal13/finalprojecthactiv8/models"
	"github.com/iqbal13/finalprojecthactiv8/structs"
	"github.com/iqbal13/finalprojecthactiv8/usecase"
)

type commentHandler struct {
	commentUseCase usecase.CommentUseCase
	photoUseCase   usecase.PhotoUseCase
}

func NewCommentHandler(routers *gin.Engine, commentUseCase usecase.CommentUseCase, photoUseCase usecase.PhotoUseCase) {
	handler := &commentHandler{commentUseCase, photoUseCase}

	router := routers.Group("/comments")
	{
		router.Use(middleware.Authentication())
		router.GET("", handler.Fetch)
		router.POST("", handler.Store)
		router.PUT("/:commentId", middleware.CommentAuthorization(handler.commentUseCase), handler.Update)
		router.DELETE("/:commentId", middleware.CommentAuthorization(handler.commentUseCase), handler.Delete)
	}
}

func (handler *commentHandler) Fetch(ctx *gin.Context) {
	var (
		comments []models.Comment

		err error
	)

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := string(userData["id"].(string))

	if err = handler.commentUseCase.Fetch(ctx.Request.Context(), &comments, userID); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "fail",
			Message: err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, helpers.ResponseData{
		Status: "success",
		Data:   comments,
	})
}

func (handler *commentHandler) Store(ctx *gin.Context) {
	var (
		comment models.Comment
		photo   models.Photo
		err     error
	)

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := string(userData["id"].(string))

	if err = ctx.ShouldBindJSON(&comment); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "fail",
			Message: err.Error(),
		})

		return
	}

	photoID := comment.PhotoID

	if err = handler.photoUseCase.GetByID(ctx.Request.Context(), &photo, photoID); err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, helpers.ResponseMessage{
			Status:  "fail",
			Message: fmt.Sprintf("photo with id %s doesn't exist", photoID),
		})

		return
	}

	comment.UserID = userID

	if err = handler.commentUseCase.Store(ctx.Request.Context(), &comment); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "fail",
			Message: err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusCreated, helpers.ResponseData{
		Status: "success",
		Data: structs.AddedComment{
			ID:        comment.ID,
			UserID:    comment.UserID,
			PhotoID:   comment.PhotoID,
			Message:   comment.Message,
			CreatedAt: comment.CreatedAt,
		},
	})
}
func (handler *commentHandler) Update(ctx *gin.Context) {
	var (
		comment models.Comment
		photo   models.Photo
		err     error
	)

	commentID := ctx.Param("commentId")
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := string(userData["id"].(string))

	if err = ctx.ShouldBindJSON(&comment); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "fail",
			Message: err.Error(),
		})

		return
	}

	updatedComment := models.Comment{
		UserID:  userID,
		Message: comment.Message,
	}

	if photo, err = handler.commentUseCase.Update(ctx.Request.Context(), updatedComment, commentID); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "fail",
			Message: err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, helpers.ResponseData{
		Status: "success",
		Data: structs.UpdatedComment{
			ID:        photo.ID,
			UserID:    photo.UserID,
			Title:     photo.Title,
			PhotoUrl:  photo.PhotoUrl,
			Caption:   photo.Caption,
			UpdatedAt: photo.UpdatedAt,
		},
	})
}
func (handler *commentHandler) Delete(ctx *gin.Context) {
	commentID := ctx.Param("commentId")

	if err := handler.commentUseCase.Delete(ctx.Request.Context(), commentID); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "fail",
			Message: err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Comment berhasil dihapus",
	})
}
