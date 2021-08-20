package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/namnguyen191/themuzix-golang-rest-api/dto"
	"github.com/namnguyen191/themuzix-golang-rest-api/entity"
	"github.com/namnguyen191/themuzix-golang-rest-api/helper"
	"github.com/namnguyen191/themuzix-golang-rest-api/service"
)

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	authService service.AuthService
	jwtService  service.JWTService
}

func (c *authController) Login(ctx *gin.Context) {
	var loginDTO dto.LoginDTO
	errDTO := ctx.ShouldBind(&loginDTO)
	if errDTO != nil {
		res := helper.BuildReponse(http.StatusBadRequest, "failed to process request", []error{errors.New("error parsing loginDTO")}, helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
	}

	authResult := c.authService.VerifyCredential(loginDTO.Email, loginDTO.Password)
	if v, ok := authResult.(entity.User); ok {
		generatedToken := c.jwtService.GenerateToken(strconv.FormatUint(uint64(v.ID), 10))
		v.Token = generatedToken
		res := helper.BuildReponse(http.StatusOK, "ok", nil, v)
		ctx.JSON(http.StatusOK, res)
		return
	}

	res := helper.BuildReponse(http.StatusUnauthorized, "invalid credential", []error{errors.New("invalid credential")}, nil)
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
}

func (c *authController) Register(ctx *gin.Context) {
	var registerDTo dto.RegisterDTO
	errorDTO := ctx.ShouldBind(&registerDTo)
	if errorDTO != nil {
		res := helper.BuildReponse(http.StatusBadRequest, "failed to process request", []error{errors.New("error parsing loginDTO")}, nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
	}

	if !c.authService.IsDuplicateEmail(registerDTo.Email) {
		res := helper.BuildReponse(http.StatusBadRequest, "duplicated email", []error{errors.New("error inserting duplicated email")}, nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
	} else {
		createdUser := c.authService.CreateUser(registerDTo)
		token := c.jwtService.GenerateToken(strconv.FormatUint(uint64(createdUser.ID), 10))
		createdUser.Token = token
		res := helper.BuildReponse(http.StatusOK, "ok", nil, createdUser)
		ctx.JSON(http.StatusOK, res)
	}
}

func NewAuthController(authService service.AuthService, jwtService service.JWTService) AuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}
