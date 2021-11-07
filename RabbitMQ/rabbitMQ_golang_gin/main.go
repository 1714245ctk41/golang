package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default() //new gin router initialization
	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"data": "Hello world!"})
	}) //* first endpoint returns hello world

	router.Run(":8000") //* running application, Default port is 8000
}
