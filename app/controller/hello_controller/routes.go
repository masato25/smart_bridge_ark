package hello_controller

import (
	"github.com/gin-gonic/gin"
)

func Route(r *gin.Engine) {
	r.GET("/ping", H1)
	r.GET("/rate", GetRates)
	r.GET("/price_rate", GetRatesHtml)
	r.GET("/", GetRatesHtml)
}
