package handlers

import (
	conte "context"
	"crawl_data/pkg/model"
	"crawl_data/pkg/service"
	"crawl_data/pkg/utils"
	"log"
	"os"

	rabbitmq_go "crawl_data/pkg/utils/rabbitmq_go"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//UserController is a ....
type VascaraController interface {
	InsertProductSendMessage(context *gin.Context)
	InsertProductReceiveMessage(context *gin.Context)
	InsertProductTest(context *gin.Context)
}

type vascaraController struct {
	vascaraService service.VascaraService
	jwtService     service.JWTService
	historyService service.HistoryService
}

//NewUserController is creating anew instance of UserControlller
func NewVascaraController(vascaraService service.VascaraService, jwtService service.JWTService, historyService service.HistoryService) VascaraController {
	return &vascaraController{
		vascaraService: vascaraService,
		jwtService:     jwtService,
		historyService: historyService,
	}
}

var linkFile = "./pkg/utils/rabbitmq_go/vascara_cache/"

func (v *vascaraController) InsertProductTest(context *gin.Context) {
	productBestSale := rabbitmq_go.BestSaleHandle()

	client := utils.GetConnection()
	defer func() {
		if err := client.Disconnect(conte.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	str := "https://www.vascara.com/giay-cao-got/giay-sandal-buoc-day-vascara-x-le-thanh-hoa-like-the-sunshine-sdn-0706-mau-trang"
	err := v.vascaraService.InsertProduct(str, client, productBestSale, new(int))
	if err != nil {
		context.JSON(400, gin.H{
			"message": "Fail insert product",
			"Error":   err,
		})
	}
	context.JSON(200, gin.H{
		"message": "Success insert product",
	})
}

func (v *vascaraController) InsertProductReceiveMessage(context *gin.Context) {
	err := os.RemoveAll(linkFile + "cache")
	if err != nil {
		log.Fatal(err)
	}
	authHeader := context.GetHeader("Authorization")
	authHeader = strings.TrimPrefix(authHeader, "Bearer ")
	token, err := v.jwtService.ValidateToken(authHeader)
	if err != nil {
		panic(err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	role := fmt.Sprintf("%v", claims["role"])

	if strings.Compare(role, "admin") == 0 {

		v.vascaraService.ReadMessage()
		context.JSON(200, gin.H{
			"message": "Reading messages",
		})
		return

	}
	context.JSON(400, gin.H{
		"message": "Error Authorization",
	})
	return

}

func (v *vascaraController) InsertProductSendMessage(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")
	authHeader = strings.TrimPrefix(authHeader, "Bearer ")
	token, err := v.jwtService.ValidateToken(authHeader)
	if err != nil {
		panic(err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	role := fmt.Sprintf("%v", claims["role"])

	if strings.Compare(role, "admin") == 0 {
		url := context.FullPath()
		if strings.Contains(url, "vascara") {
			hitory := model.History{
				Url: url,
			}
			err = v.historyService.InsertHistory(hitory)
			if err != nil {
				res := utils.BuildErrorResponse("Error InsertHistory: ", err.Error(), nil)
				context.JSON(http.StatusBadRequest, res)
				return

			}
			productReview := []model.ProductView{}
			productReview, err := v.vascaraService.GetProductReview()
			if err != nil {
				context.JSON(300, gin.H{
					"message": "Error get productview",
				})
				return
			}

			rabbitmq_go.Message(productReview)

			res := utils.BuildResponse(true, "OK!", nil)
			context.JSON(http.StatusOK, res)
			return

		} else {
			res := utils.BuildErrorResponse("Error ignore string: ", err.Error(), nil)
			context.JSON(http.StatusBadRequest, res)
			return

		}

	} else {
		// res := utils.BuildErrorResponse("Error Authotization: ", err.Error(), nil)
		// context.JSON(http.StatusBadRequest, res)
		context.JSON(400, gin.H{
			"error": "You don't have permission",
		})
		return
	}

}
