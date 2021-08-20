package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
		res := helper.BuildResponse(
			http.StatusBadRequest,
			"failed to process request",
			helper.ValidationErrorsToStringArray(errDTO.(validator.ValidationErrors)),
			nil,
		)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	authResult := c.authService.VerifyCredential(loginDTO.Email, loginDTO.Password)
	if v, ok := authResult.(entity.User); ok {
		generatedToken := c.jwtService.GenerateToken(strconv.FormatUint(uint64(v.ID), 10))
		res := helper.BuildResponse(http.StatusOK, "ok", nil, &dto.LoginResponseDTO{
			Name:  v.Name,
			Email: v.Email,
			Token: generatedToken,
		})
		ctx.JSON(http.StatusOK, res)
		return
	}

	res := helper.BuildResponse(
		http.StatusUnauthorized,
		"invalid credential", []string{"invalid credential"}, nil)
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
}

func (c *authController) Register(ctx *gin.Context) {
	var registerDTo dto.RegisterDTO
	errorDTO := ctx.ShouldBind(&registerDTo)
	if errorDTO != nil {
		res := helper.BuildResponse(
			http.StatusBadRequest,
			"failed to process request",
			helper.ValidationErrorsToStringArray(errorDTO.(validator.ValidationErrors)),
			nil,
		)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	if !c.authService.IsDuplicateEmail(registerDTo.Email) {
		res := helper.BuildResponse(
			http.StatusBadRequest,
			"duplicated email",
			[]string{"email has already been used"},
			nil,
		)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
	} else {
		createdUser := c.authService.CreateUser(registerDTo)
		token := c.jwtService.GenerateToken(strconv.FormatUint(uint64(createdUser.ID), 10))
		createdUser.Token = token
		res := helper.BuildResponse(http.StatusOK, "ok", nil, createdUser)
		ctx.JSON(http.StatusOK, res)
	}
}

func NewAuthController(authService service.AuthService, jwtService service.JWTService) AuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}
