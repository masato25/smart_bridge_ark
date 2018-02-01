package arklib

import (
	"github.com/masato25/ark-go/core"
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
