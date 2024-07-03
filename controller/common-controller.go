package controller

import (
	"github.com/gin-gonic/gin"
	service "github.com/richardgong1987/smartcloud-prices/services"
	"net/http"
)

type ControllerPrices struct {
}

var servicePrice = &service.SmartcloudPriceService{}

func (h *ControllerPrices) GetPricesHandler(c *gin.Context) {
	kind := c.Query("kind")
	if kind == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "kind parameter is required"})
		return
	}

	prices, err := servicePrice.GetPrices(kind)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch prices"})
		return
	}

	c.JSON(http.StatusOK, prices)
}
