package middlewares

import (
	"net/http"

	"github.com/chentihe/mongodb-api/utils"
	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware(publicKey string) func(c *gin.Context) {
	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")
		token, err := utils.ExtractToken(header)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}

		err = utils.ValidateToken(token, publicKey)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}

		c.Next()
	}
}
