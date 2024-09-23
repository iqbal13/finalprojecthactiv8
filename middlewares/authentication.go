package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iqbal13/finalprojecthactiv8/helpers"
)

func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		verifyToken, err := helpers.VerifyToken(ctx)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, helpers.ResponseMessage{
				Status:  "Unauthorized",
				Message: err.Error(),
			})

			return
		}

		ctx.Set("userData", verifyToken)
		ctx.Next()
	}
}
