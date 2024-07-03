package main

import (
	"github.com/gin-gonic/gin"
	controller "github.com/richardgong1987/smartcloud-prices/controller"
	"log"
)

func main() {
	controllerPrice := &controller.ControllerPrices{}
	router := gin.Default()
	router.GET("/prices", controllerPrice.GetPricesHandler)
	log.Println("Server is running on port 8080")
	log.Fatal(router.Run(":8080"))
}
