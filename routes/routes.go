package routes

import (
	"github.com/chentihe/mongodb-api/config/svc"
	"github.com/chentihe/mongodb-api/middlewares"
	"github.com/gin-gonic/gin"
)

//	@title			News API
//	@version		1.0
//	@description	This is an server to manage news from mongoDB.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/api/v1

//	@securityDefinitions.apikey	BearerAuth
//	@in							header
//	@name						Authorization
//	@description				Type "Bearer" followed by a space and JWT token

//	@externalDocs.description	OpenAPI
//	@externalDocs.url			https://swagger.io/resources/open-api/
func RegisterRouters(router *gin.Engine, ctx *svc.ServiceContext) {
	v1 := router.Group("/api/v1")
	AddAuthRoutes(v1, ctx)
	v1.Use(middlewares.JWTAuthMiddleware(ctx.Config.Jwt.PublicKey))
	AddMediaRoutes(v1, ctx)
}
