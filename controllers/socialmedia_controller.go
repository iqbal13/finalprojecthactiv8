package controllers

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/iqbal13/finalprojecthactiv8/helpers"
	middleware "github.com/iqbal13/finalprojecthactiv8/middlewares"
	"github.com/iqbal13/finalprojecthactiv8/models"
	"github.com/iqbal13/finalprojecthactiv8/structs"
	"github.com/iqbal13/finalprojecthactiv8/usecase"
)

type socialMediaHandler struct {
	socialMediaUseCase usecase.SocialMediaUseCase
}

func NewSocialMediaHandler(routers *gin.Engine, socialMediaUseCase usecase.SocialMediaUseCase) {
	handler := &socialMediaHandler{socialMediaUseCase}

	router := routers.Group("/socialmedia")
	{
		router.Use(middleware.Authentication())
		router.GET("", handler.Fetch)
		router.POST("", handler.Store)
		router.PUT("/:socialMediaId", middleware.SosmedAuthorization(handler.socialMediaUseCase), handler.Update)
		router.DELETE("/:socialMediaId", middleware.SosmedAuthorization(handler.socialMediaUseCase), handler.Delete)
	}
}

func (handler *socialMediaHandler) Fetch(ctx *gin.Context) {
	var (
		socialMedias []models.SocialMedia
		err          error
	)

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := string(userData["id"].(string))

	if err = handler.socialMediaUseCase.Fetch(ctx.Request.Context(), &socialMedias, userID); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "fail",
			Message: err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, helpers.ResponseData{
		Status: "success",
		Data: structs.FetchedSocialMedia{
			SocialMedias: socialMedias,
		},
	})
}
func (handler *socialMediaHandler) Store(ctx *gin.Context) {
	var (
		socialMedia models.SocialMedia
		err         error
	)

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := string(userData["id"].(string))

	if err = ctx.ShouldBindJSON(&socialMedia); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "fail",
			Message: err.Error(),
		})

		return
	}

	socialMedia.UserID = userID

	if err = handler.socialMediaUseCase.Store(ctx.Request.Context(), &socialMedia); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "fail",
			Message: err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusCreated, helpers.ResponseData{
		Status: "success",
		Data: structs.AddedSocialMedia{
			ID:             socialMedia.ID,
			UserID:         socialMedia.UserID,
			Name:           socialMedia.Name,
			SocialMediaUrl: socialMedia.SocialMediaUrl,
			CreatedAt:      socialMedia.CreatedAt,
		},
	})
}

func (handler *socialMediaHandler) Update(ctx *gin.Context) {
	var (
		socialMedia models.SocialMedia
		err         error
	)

	socialMediaID := ctx.Param("socialMediaId")
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := string(userData["id"].(string))

	if err = ctx.ShouldBindJSON(&socialMedia); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "fail",
			Message: err.Error(),
		})

		return
	}

	updatedSocialMedia := models.SocialMedia{
		UserID:         userID,
		Name:           socialMedia.Name,
		SocialMediaUrl: socialMedia.SocialMediaUrl,
	}

	if socialMedia, err = handler.socialMediaUseCase.Update(ctx.Request.Context(), updatedSocialMedia, socialMediaID); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "fail",
			Message: err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, structs.UpdatedSocialMedia{
		ID:             socialMedia.ID,
		Name:           socialMedia.Name,
		SocialMediaUrl: socialMedia.SocialMediaUrl,
		UserID:         socialMedia.UserID,
		UpdatedAt:      socialMedia.UpdatedAt,
	})
}

func (handler *socialMediaHandler) Delete(ctx *gin.Context) {
	socialMediaID := ctx.Param("socialMediaId")

	if err := handler.socialMediaUseCase.Delete(ctx.Request.Context(), socialMediaID); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "fail",
			Message: err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Social Media berhasil dihapus",
	})
}
