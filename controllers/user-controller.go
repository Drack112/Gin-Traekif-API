package controllers

import (
    "fmt"
    "log"
    "net/http"
    "strconv"

    "github.com/Drack112/Crud-Golang-API/dto"
    "github.com/Drack112/Crud-Golang-API/helper"
    "github.com/Drack112/Crud-Golang-API/service"
    "github.com/dgrijalva/jwt-go"
    "github.com/gin-gonic/gin"
)

type UserController interface {
    Update(ctx *gin.Context)
    Profile(ctx *gin.Context)
}

type userController struct {
    userService service.UserService
    jwtService  service.JWTService
}

func NewUserController(userService service.UserService, jwtService service.JWTService) UserController {
    return &userController{
        userService: userService,
        jwtService:  jwtService,
    }
}

func (controllerUser *userController) Update(ctx *gin.Context) {
    var userUpdateDTO dto.UserUpdateDTO
    errDTO := ctx.ShouldBind(userUpdateDTO)
    if errDTO != nil {
        response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
        ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
        return
    }

    authHeader := ctx.GetHeader("Authorization")
    token, errToken := controllerUser.jwtService.ValidateToken(authHeader)
    if errToken != nil {
        log.Panicf("Error in validate the JWT Token\n %s", errToken.Error())
    }
    claims := token.Claims.(jwt.MapClaims)

    id, err := strconv.ParseUint(fmt.Sprintf("%v", claims["user_id"]), 10, 64)
    if err != nil {
        log.Panicf("Error in parsing the Claims of Token\n %s", errToken.Error())
    }

    userUpdateDTO.ID = id
    u := controllerUser.userService.Update(userUpdateDTO)
    res := helper.BuildResponse(true, "Ok!", u)
    ctx.JSON(http.StatusOK, res)
}

func (controllerUser *userController) Profile(ctx *gin.Context) {
    authHeader := ctx.GetHeader("Authorization")
    token, errToken := controllerUser.jwtService.ValidateToken(authHeader)
    if errToken != nil {
        log.Panicf("Error in validate the JWT Token\n %s", errToken.Error())
    }

    claims := token.Claims.(jwt.MapClaims)
    id := fmt.Sprintf("%v", claims["user_id"])
    user := controllerUser.userService.Profile(id)
    res := helper.BuildResponse(true, "Ok!", user)
    ctx.JSON(http.StatusOK, res)

}
