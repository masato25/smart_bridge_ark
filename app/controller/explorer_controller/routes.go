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
	r.GET("/arks/update_voters", UpdateVoterController)
	r.GET("/action/arks/sync_blocks", SyncNewBlocks)
	r.GET("/action/arks/sync_block_with_time", GetBlockInfoByTimeLimit)
	r.GET("/action/arks/count_rewards", CalculatorEachBlocks)

	rapi := r.Group("/api/v1")
	rapi.GET("/data/voters.json", GetVoterController)
	rapi.GET("/data/arks.json", ArkTransactionsJSON)
	rapi.GET("/data/eths.json", EthTransactionsJSON)
	rapi.GET("/data/arks/number_of_blocks.json", GetCreatedBlockNumber)
	rapi.GET("/data/profit/:address", QueryVoteProfit)

	// views
	r.GET("/votes", VotesView)
	r.GET("/votes/:address", VotesProfitView)
}
