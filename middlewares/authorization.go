package middleware

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/iqbal13/finalprojecthactiv8/helpers"
	"github.com/iqbal13/finalprojecthactiv8/models"
	"github.com/iqbal13/finalprojecthactiv8/usecase"
)

func PhotoAuthorization(photoUseCase usecase.PhotoUseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			photo models.Photo
			err   error
		)

		photoID := ctx.Param("photoId")
		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userID := string(userData["id"].(string))

		if err = photoUseCase.GetByID(ctx.Request.Context(), &photo, photoID); err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, helpers.ResponseMessage{
				Status:  "fail",
				Message: fmt.Sprintf("Photo dengan id %s tidak ditemukan!", photoID),
			})

			return
		}

		if photo.UserID != userID {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, helpers.ResponseMessage{
				Status:  "Unauthorized",
				Message: "maaf anda tidak memiliki hak untuk mengedit/menghapus photo ini",
			})

			return
		}
	}
}

func SosmedAuthorization(socialMediaUseCase usecase.SocialMediaUseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			socialMedia models.SocialMedia
			err         error
		)

		socialMediaID := ctx.Param("socialMediaId")
		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userID := string(userData["id"].(string))

		if err = socialMediaUseCase.GetByID(ctx.Request.Context(), &socialMedia, socialMediaID); err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, helpers.ResponseMessage{
				Status:  "fail",
				Message: fmt.Sprintf("social media dengan id %s tidak ditemukan", socialMediaID),
			})

			return
		}

		if socialMedia.UserID != userID {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, helpers.ResponseMessage{
				Status:  "Unauthorized",
				Message: "maaf anda tidak memiliki hak untuk mengedit/menghapus social media ini",
			})

			return
		}
	}
}

func CommentAuthorization(commentUseCase usecase.CommentUseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			comment models.Comment
			err     error
		)

		commentID := ctx.Param("commentId")
		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userID := string(userData["id"].(string))

		if err = commentUseCase.GetByID(ctx.Request.Context(), &comment, commentID); err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, helpers.ResponseMessage{
				Status:  "fail",
				Message: fmt.Sprintf("Comment dengan id %s tidak ditemukan", commentID),
			})

			return
		}

		if comment.UserID != userID {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, helpers.ResponseMessage{
				Status:  "Unauthorized",
				Message: "maaf anda tidak memiliki hak untuk mengedit/menghapus comment ini",
			})

			return
		}
	}
}
