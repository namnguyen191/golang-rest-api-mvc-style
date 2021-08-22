package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/namnguyen191/themuzix-golang-rest-api/dto"
	"github.com/namnguyen191/themuzix-golang-rest-api/helper"
	"github.com/namnguyen191/themuzix-golang-rest-api/service"
)

type UserController interface {
	Update(context *gin.Context)
	Profile(context *gin.Context)
}

type userController struct {
	userService service.UserService
	jwtService  service.JWTService
}

func (c *userController) Update(context *gin.Context) {
	var userUpdateDTO dto.UserUpdateDTO
	errDTO := context.ShouldBind((&userUpdateDTO))
	if errDTO != nil {
		errRes, ok := errDTO.(validator.ValidationErrors)
		if !ok {
			res := helper.BuildResponse(http.StatusBadRequest, "invalid request body", []string{errDTO.Error()}, nil)
			context.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
		res := helper.BuildResponse(http.StatusBadRequest, "invalid request body", helper.ValidationErrorsToStringArray(errRes), nil)
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		res := helper.BuildResponse(http.StatusUnauthorized, "bad token", []string{errToken.Error()}, nil)
		context.AbortWithStatusJSON(http.StatusUnauthorized, res)
		return
	}

	claims := token.Claims.(jwt.MapClaims)
	id, err := strconv.ParseUint(fmt.Sprintf("%v", claims["user_id"]), 10, 64)
	if err != nil {
		panic(err.Error())
	}
	userUpdateDTO.ID = uint(id)
	u := c.userService.Update(userUpdateDTO)
	res := helper.BuildResponse(http.StatusOK, "ok", nil, u)
	context.JSON(http.StatusOK, res)
}

func (c *userController) Profile(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")
	token, err := c.jwtService.ValidateToken(authHeader)
	if err != nil {
		res := helper.BuildResponse(http.StatusUnauthorized, "bad token", []string{err.Error()}, nil)
		context.AbortWithStatusJSON(http.StatusUnauthorized, res)
		return
	}

	claims := token.Claims.(jwt.MapClaims)
	user := c.userService.Profile(fmt.Sprintf("%v", claims["user_id"]))
	res := helper.BuildResponse(http.StatusOK, "ok", nil, user)
	context.JSON(http.StatusOK, res)
}

func NewUserController(userService service.UserService, jwtService service.JWTService) UserController {
	return &userController{
		userService: userService,
		jwtService:  jwtService,
	}
}
