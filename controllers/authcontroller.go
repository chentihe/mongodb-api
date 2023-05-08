package controllers

import (
	"net/http"

	"github.com/chentihe/mongodb-api/config"
	"github.com/chentihe/mongodb-api/dtos"
	"github.com/chentihe/mongodb-api/services"
	"github.com/chentihe/mongodb-api/utils"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	UserService *services.UserService
	Jwt         *config.Jwt
}

func NewAuthController(userService *services.UserService, jwt *config.Jwt) *AuthController {
	return &AuthController{
		UserService: userService,
		Jwt:         jwt,
	}
}

// AuthHandler @Summary
//
//	@Tags		user
//	@Version	1.0
//	@Produce	application/json
//	@Param		register	body	dtos.LoginDto	true	"login"
//	@Success	200			string	successful		return	token
//	@Router		/login [post]
func (c *AuthController) Login(ctx *gin.Context) {
	var loginDto dtos.LoginDto
	if err := ctx.ShouldBindJSON(&loginDto); err != nil {
		panic(err)
	}

	user, err := c.UserService.GetUserByName(loginDto.UserName)
	if err != nil {
		panic(err)
	}

	if err = utils.VerifyPassword(user.Password, loginDto.Password); err != nil {
		panic(err)
	}

	accessToken, err := utils.GenerateToken(c.Jwt.ExpiresIn, user.Id, c.Jwt.PrivateKey)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "access_token": accessToken})
}
