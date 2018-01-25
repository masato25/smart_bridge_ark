package explorer_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/masato25/smart_bridge_ark/app/model/ark"
	_ "github.com/masato25/smart_bridge_ark/app/model/ether"
)

func ArkTransactions(c *gin.Context) {
	arks := []ark.ArkTransaction{}
	db.Find(&arks)
	c.JSON(200, gin.H{
		"data": arks,
	})
	return
}
