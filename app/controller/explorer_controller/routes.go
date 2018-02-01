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
	r.GET("/data/eths", EthTransactions)
	r.GET("/action/ark/sync", SyncArkTranactions)
	r.GET("/action/eth/sync", SyncEthTranactions)

	rapi := r.Group("/api/v1")
	rapi.GET("/data/arks.json", ArkTransactionsJSON)
	rapi.GET("/data/eths.json", EthTransactionsJSON)
}
