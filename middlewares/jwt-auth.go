package middlewares

import (
    "log"
    "net/http"

    "github.com/Drack112/Crud-Golang-API/helper"
    "github.com/Drack112/Crud-Golang-API/service"
    "github.com/dgrijalva/jwt-go"
    "github.com/gin-gonic/gin"
)

func AuthorizeJWT(jwtService service.JWTService) gin.HandlerFunc {
    return func(ctx *gin.Context) {
        authHeader := ctx.GetHeader("Authorization")

        if authHeader == "" {
            response := helper.BuildErrorResponse("Failed to process request", "No token found", nil)
            ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
        }

        token, err := jwtService.ValidateToken(authHeader)

        if token.Valid {
            claims := token.Claims.(jwt.MapClaims)
            log.Println("Claim[user_id]: ", claims["user_id"])
            log.Println("Claim[issuer]: ", claims["issuer"])
        } else {
            log.Println(err)
            response := helper.BuildErrorResponse("Token is not valid", err.Error(), nil)
            ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
        }
    }
}
