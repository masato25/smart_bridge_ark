package arklib

import (
	"github.com/labstack/gommon/log"
	"github.com/masato25/ark-go/core"
	"github.com/masato25/smart_bridge_ark/config"
)

func GetTransactionsSender(senderId string) (resp *[]core.Transaction) {
	params := core.TransactionQueryParams{
		SenderID: senderId,
		OrderBy:  "timestamp:desc",
	}
	httpresp, _, _ := arkclient.ListTransaction(params)
	resp = &httpresp.Transactions
	return
}

func GetTransactionsReceive(recipeint string) (resp *[]core.Transaction) {
	params := core.TransactionQueryParams{
		RecipientID: recipeint,
		OrderBy:     "timestamp:desc",
	}
	httpresp, _, _ := arkclient.ListTransaction(params)
	resp = &httpresp.Transactions
	return
}

func GetBlockGenerated() core.BlocksResponse {
	bodyrespon, arkresp, _ := arkclient.GetBlocksByGeneratorPublicKey(config.MyConfig().Ark.PublicKey, 100)
	if arkresp.Error() != "" {
		log.Error(arkresp.Error())
	}
	// log.Info("a: %v, b: %v, c: %v", httpresp, b, c)
	return bodyrespon
}

func GetBlockGeneratedByHeight(height int) core.BlocksResponse {
	bodyrespon, arkresp, _ := arkclient.GetBlocksByHeight(height)
	if arkresp.Error() != "" {
		log.Error(arkresp.Error())
	}
	// log.Info("a: %v, b: %v, c: %v", httpresp, b, c)
	return bodyrespon
}

func GetCurrentHeight() int {
	currentHeight, _, _ := arkclient.GetPeerHeight()
	fheight := currentHeight.Height
	return fheight
}
