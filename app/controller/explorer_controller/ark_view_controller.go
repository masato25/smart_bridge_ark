package explorer_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/masato25/smart_bridge_ark/app/model/ether"
)

func VotesView(c *gin.Context) {
	c.HTML(http.StatusOK, "voters_index", gin.H{})
	return
}

func VotesProfitView(c *gin.Context) {
	c.HTML(http.StatusOK, "voters_profit_index", gin.H{})
	return
}
