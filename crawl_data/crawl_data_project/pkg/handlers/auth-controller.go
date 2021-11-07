package handlers

import (
	"crawl_data/pkg/model"
	"crawl_data/pkg/service"
	"crawl_data/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

//AuthController interface is a contract what this controller can do
type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	authService service.AuthService
	jwtService  service.JWTService
}

//NewAuthController creates a new instance of AuthController
func NewAuthController(authService service.AuthService, jwtService service.JWTService) AuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (c *authController) Login(ctx *gin.Context) {
	loginDTO := &model.LoginDTO{}
	errDTO := ctx.ShouldBind(loginDTO)
	if errDTO != nil {
		response := utils.BuildErrorResponse("Failed to process request", errDTO.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authResult := c.authService.VerifyCredential(loginDTO.Email, loginDTO.Password)
	if v, ok := authResult.(model.User); ok {
		generatedToken := c.jwtService.GenerateToken(v.ID, v.Role)
		v.Token = generatedToken
		response := utils.BuildResponse(true, "OK!", v)
		ctx.JSON(http.StatusOK, response)
		return
	}
	response := utils.BuildErrorResponse("Please check again your credential", "Invalid Credential", utils.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

func (c *authController) Register(ctx *gin.Context) {

	registerDTO := &model.RegisterDTO{}
	errDTO := ctx.ShouldBind(registerDTO)
	if errDTO != nil {
		response := utils.BuildErrorResponse("Failed to process request", errDTO.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if !c.authService.IsDuplicateEmail(registerDTO.Email) {
		response := utils.BuildErrorResponse("Failed to process request", "Duplicate email", utils.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
	}
	createdUser, err := c.authService.CreateUser(registerDTO)
	if err != nil {
		res := utils.BuildErrorResponse("Error update user", err.Error(), createdUser)
		ctx.JSON(http.StatusBadRequest, res)
	}
	token := c.jwtService.GenerateToken(createdUser.ID, createdUser.Role)

	createdUser.Token = token
	response := utils.BuildResponse(true, "OK!", createdUser)
	ctx.JSON(http.StatusCreated, response)

}
