package explorer_controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/masato25/ark-go/core"
	"github.com/masato25/smart_bridge_ark/app/model/ark"
	_ "github.com/masato25/smart_bridge_ark/app/model/ether"
	"github.com/masato25/smart_bridge_ark/config"
	"github.com/masato25/smart_bridge_ark/lib/arklib"
)

func ArkTransactions(c *gin.Context) {
	c.HTML(http.StatusOK, "ark_index", gin.H{})
	return
}

func ArkTransactionsJSON(c *gin.Context) {
	arks := []ark.ArkTransaction{}
	db.Find(&arks)
	c.JSON(http.StatusOK, gin.H{
		"data": arks,
	})
	return
}

func SyncArkTranactions(c *gin.Context) {
	addr := config.MyConfig().Ark.Address
	transactions := arklib.GetTransactionsReceive(addr)
	if transactions != nil {
		for _, ts := range *transactions {
			transactionTime := time.Unix(core.GetTransactionTime(ts.Timestamp).Unix(), 0)
			newRecord := ark.ArkTransaction{
				ID:              ts.ID,
				FromAddr:        ts.SenderID,
				ToAddr:          ts.RecipientID,
				BlockId:         ts.Blockid,
				Signature:       ts.Signature,
				SenderPublicKey: ts.SenderPublicKey,
				TransactionTime: transactionTime,
				SmartBridge:     ts.VendorField,
				Amount:          arklib.ToArk(ts.Amount),
				Fee:             ts.Fee,
			}

			db.Create(&newRecord)
		}
		c.JSON(200, gin.H{
			"data": *transactions,
		})
	} else {
		c.JSON(200, gin.H{
			"data": []string{},
		})
	}

	return
}
