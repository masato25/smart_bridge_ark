package explorer_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/masato25/smart_bridge_ark/app/setup"
)

var db *gorm.DB

func Route(r *gin.Engine) {
	db = setup.GetConn()
	r.GET("/data/arks", ArkTransactions)
}
