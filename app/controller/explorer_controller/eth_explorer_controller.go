package explorer_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/masato25/smart_bridge_ark/app/model/ark"
	"github.com/masato25/smart_bridge_ark/app/model/ether"
	"github.com/masato25/smart_bridge_ark/lib/ethereumrpc"
	log "github.com/sirupsen/logrus"
)

func EthTransactions(c *gin.Context) {
	c.HTML(http.StatusOK, "eth_index", gin.H{})
	return
}

func EthTransactionsJSON(c *gin.Context) {
	ethers := []ether.EtherTransaction{}
	db.Find(&ethers)
	c.JSON(http.StatusOK, gin.H{
		"data": ethers,
	})
	return
}

func SyncEthTranactions(c *gin.Context) {
	arks := []ark.ArkTransaction{}
	db.Find(&arks)
	var errors []string
	for _, ak := range arks {
		err := ethereumrpc.SyncSendEth(ak)
		if err != nil {
			errors = append(errors, err.Error())
			log.Error(err)
		}
	}
	if len(errors) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	} else {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"errors": errors,
		})
	}
	return
}
