package handlers

import (
	"crawl_data/pkg/model"
	"crawl_data/pkg/service"
	"crawl_data/pkg/utils"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//UserController is a ....
type UserController interface {
	Update(context *gin.Context)
	Profile(context *gin.Context)
}

type userController struct {
	userService service.UserService
	jwtService  service.JWTService
	authService service.AuthService
}

//NewUserController is creating anew instance of UserControlller
func NewUserController(userService service.UserService, jwtService service.JWTService, authService service.AuthService) UserController {
	return &userController{
		userService: userService,
		jwtService:  jwtService,
		authService: authService,
	}
}

func (c *userController) Update(context *gin.Context) {
	userUpdateDTO := &model.UserUpdateDTO{}
	errDTO := context.ShouldBind(userUpdateDTO)
	if errDTO != nil {
		res := utils.BuildErrorResponse("Failed to process request", errDTO.Error(), utils.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	authHeader := context.GetHeader("Authorization")
	authHeader = strings.TrimPrefix(authHeader, "Bearer ")

	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	idUser := fmt.Sprintf("%v", claims["user_id"])

	userUpdateDTO.ID = idUser

	u, err := c.userService.Update(userUpdateDTO)
	if err != nil {
		res := utils.BuildErrorResponse("Error update user", err.Error(), u)
		context.JSON(http.StatusBadRequest, res)
	}
	tokenNew := c.jwtService.GenerateToken(u.ID, u.Role)
	u.Token = tokenNew

	res := utils.BuildResponse(true, "OK!", u)
	context.JSON(http.StatusOK, res)

}

func (c *userController) Profile(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")
	authHeader = strings.TrimPrefix(authHeader, "Bearer ")
	token, err := c.jwtService.ValidateToken(authHeader)
	if err != nil {
		panic(err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	user, err := c.userService.Profile(id)
	if err != nil {
		res := utils.BuildErrorResponse("Error update user", err.Error(), user)
		context.JSON(http.StatusBadRequest, res)
	}
	res := utils.BuildResponse(true, "OK", user)
	context.JSON(http.StatusOK, res)

}
