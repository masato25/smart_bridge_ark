package explorer_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/masato25/smart_bridge_ark/app/model/delegate"
	_ "github.com/masato25/smart_bridge_ark/app/model/ether"
)

func QueryVoteProfit(c *gin.Context) {
	address := c.Param("address")
	qProfit := delegate.VoteProfit{
		VoteID: address,
	}
	var qProfits []delegate.VoteProfit
	db.Where(&qProfit).Find(&qProfits)
	var sumReward float64
	db.Where(&qProfit).Table(qProfit.TableName()).Select("sum(reward)").Row().Scan(&sumReward)
	c.JSON(200, gin.H{
		"data": gin.H{
			"reward_per_blocks": qProfits,
			"reward_sum":        sumReward,
		},
	})
	return
}
