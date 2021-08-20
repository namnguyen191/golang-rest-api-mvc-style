package middleware

import (
	"errors"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/namnguyen191/themuzix-golang-rest-api/helper"
	"github.com/namnguyen191/themuzix-golang-rest-api/service"
)

func AuthorizeJWT(jwtServ service.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response := helper.BuildReponse(http.StatusUnauthorized, "missing auth header", []error{errors.New("missing auth header")}, nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		token, err := jwtServ.ValidateToken(authHeader)

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claim[user_id]: ", claims["user_id"])
			log.Println("Claim[issuer]: ", claims["issuer"])
		} else {
			log.Println(err)
			response := helper.BuildReponse(http.StatusUnauthorized, "invalid token", []error{errors.New("invalid token")}, nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
	}
}
