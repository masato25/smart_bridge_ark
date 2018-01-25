package hello_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/masato25/smart_bridge_ark/lib/exchange_rate"
)

func H1(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
	return
}

func GetRates(c *gin.Context) {
	rates := exchange_rate.GetAllRate()
	c.JSON(200, gin.H{
		"data": rates,
	})
	return
}

func GetRatesHtml(c *gin.Context) {
	rates := exchange_rate.GetAllRate()
	c.HTML(http.StatusOK, "crypto_index", gin.H{
		"title": "Price Rate",
		"data":  rates,
	})
	return
}
