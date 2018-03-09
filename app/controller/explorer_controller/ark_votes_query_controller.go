package explorer_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/masato25/smart_bridge_ark/app/model/delegate"
	_ "github.com/masato25/smart_bridge_ark/app/model/ether"
)

type QueryVoteProfitInput struct {
	Showpaid bool `json:"showpaid" form:"showpaid"`
}

func QueryVoteProfit(c *gin.Context) {
	var inputs QueryVoteProfitInput
	if err := c.Bind(&inputs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	address := c.Param("address")
	qProfit := delegate.VoteProfit{
		VoteID: address,
	}
	var qProfits []delegate.VoteProfit
	var sumReward float64
	if inputs.Showpaid {
		db.Where(&qProfit).Find(&qProfits)
		db.Where(&qProfit).Table(qProfit.TableName()).Select("sum(reward)").Row().Scan(&sumReward)
	} else {
		qProfit.Payment = false
		db.Where("vote_id = ? AND payment = ?", qProfit.VoteID, qProfit.Payment).Find(&qProfits)
		db.Where("vote_id = ? AND payment = ?", qProfit.VoteID, qProfit.Payment).Table(qProfit.TableName()).Select("sum(reward)").Row().Scan(&sumReward)
	}
	c.JSON(200, gin.H{
		"data": gin.H{
			"reward_per_blocks": qProfits,
			"reward_sum":        sumReward,
		},
	})
	return
}
