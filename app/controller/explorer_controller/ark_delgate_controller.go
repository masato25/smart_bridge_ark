package explorer_controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/masato25/ark-go/core"
	"github.com/masato25/smart_bridge_ark/app/model/blocks"
	"github.com/masato25/smart_bridge_ark/app/model/delegate"
	_ "github.com/masato25/smart_bridge_ark/app/model/ether"
	"github.com/masato25/smart_bridge_ark/config"
	"github.com/masato25/smart_bridge_ark/lib/arklib"
	log "github.com/sirupsen/logrus"
)

func UpdateVoterController(c *gin.Context) {
	voter := arklib.GetVotersAndWeight()
	collectAddress := []string{}
	for _, vot := range voter {
		Dvote := delegate.Vote{
			ID:      vot.Address,
			Address: vot.Address,
			Status:  true,
			Balance: vot.Balance,
		}
		db.Where("address = ?", Dvote.Address).Save(&Dvote)
		collectAddress = append(collectAddress, Dvote.Address)
	}
	var selectUnVoter []delegate.Vote
	// get unvote list
	db.Where("address not in (?)", collectAddress).Find(&selectUnVoter)
	for _, vot := range selectUnVoter {
		db.Where("address = ?", vot.Address).Assign("status", false).FirstOrCreate(&vot)
	}
	c.JSON(200, gin.H{
		"data": voter,
	})
	return
}

type GetVoterControllerOuput struct {
	Address string  `json:"address"`
	Balance float64 `json:"balance"`
	Status  string  `json:"status"`
	Weight  float64 `json:"weight"`
}

type GetVoterControllerInput struct {
	Filter bool `json:"filter" form:"filter"`
}

func GetVoterController(c *gin.Context) {
	var input GetVoterControllerInput
	if err := c.Bind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	log.Debugf("[GetVoterController] filter: %v", input.Filter)
	Dvote := []delegate.Vote{}
	if !input.Filter {
		db.Where("status = ?", true).Find(&Dvote)
	} else {
		db.Find(&Dvote)
	}
	output := make([]GetVoterControllerOuput, len(Dvote))
	var totallBalance float64
	for _, acct := range Dvote {
		totallBalance += acct.Balance
	}
	for indx, acct := range Dvote {
		Arkbalance := acct.Balance
		output[indx].Address = acct.Address
		output[indx].Balance = Arkbalance
		output[indx].Status = fmt.Sprintf("%v", acct.Status)
		output[indx].Weight = float64((Arkbalance / totallBalance) * 100)
	}
	c.JSON(200, gin.H{
		"data": output,
	})
	return
}

func SyncNewBlocks(c *gin.Context) {
	resp := arklib.GetBlockGenerated()
	for _, block := range resp.Blocks {
		thisTime := core.GetTransactionTime(int32(block.Timestamp))
		newBlock := blocks.Block{
			CreatedAt: thisTime,
			ID:        block.ID,
			Height:    block.Height,
			Reward:    int64(block.Reward),
		}
		db.Create(&newBlock)
	}
	c.JSON(200, gin.H{
		"data": resp,
	})
	return
}

func GetCreatedBlockNumber(c *gin.Context) {
	// var blockQhelp blocks.Block
	blocks := []blocks.Block{}
	var total uint
	dt := db.Find(&blocks).Count(&total)
	if dt.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": dt.Error.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"number": total,
		})
	}
	return
}

type GetBlockInfoByTimeLimitInput struct {
	TimestampLimit int64 `json:"timestamp" form:"timestamp"`
	Fheight        int   `json:"height" form:"height"`
}

func GetBlockInfoByTimeLimit(c *gin.Context) {
	input := GetBlockInfoByTimeLimitInput{
		1520056906,
		// default should be currenct height
		arklib.GetCurrentHeight(),
	}
	if err := c.Bind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	fheight := input.Fheight
	var timestampLimit int64
	timestampLimit = input.TimestampLimit
	flag := false
	respBody := []blocks.Block{}
	for {
		resp := arklib.GetBlockGeneratedByHeight(fheight)
		for _, block := range resp.Blocks {
			thisTime := core.GetTransactionTime(int32(block.Timestamp))
			if block.GeneratorPublicKey == config.MyConfig().Ark.PublicKey {
				newBlock := blocks.Block{
					CreatedAt: thisTime,
					ID:        block.ID,
					Height:    block.Height,
					Reward:    int64(block.Reward),
				}
				db.Create(&newBlock)
				respBody = append(respBody, newBlock)
			}
			log.Infof("thisTime: %v", thisTime)
			if thisTime.Unix() < timestampLimit {
				flag = true
				break
			}
		}
		if flag {
			break
		}
		fheight -= 1
	}
	c.JSON(200, gin.H{
		"data": respBody,
	})
	return
}

type CalculatorEachBlockstInput struct {
	TimestampLimit int64 `json:"timestamp" form:"timestamp"`
}

func CalculatorEachBlocks(c *gin.Context) {
	var input CalculatorEachBlockstInput
	if err := c.Bind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	blockResults := []blocks.Block{}
	if input.TimestampLimit == 0 {
		db.Find(&blockResults)
	} else {
		etime := time.Unix(input.TimestampLimit, 0)
		db.Where("updated_at >= ?", etime).Find(&blockResults)
	}
	for _, bot := range blockResults {
		// btime := bot.CreatedAt
		votes := []delegate.Vote{}
		// db.Where("updated_at <= ? and status = ?", btime, true).Find(&votes)
		db.Where("status = ?", true).Find(&votes)
		voterList := arklib.CalculatorData(votes)
		sharedReward := float64(bot.Reward) * 0.9 / float64(100)
		for _, vot := range voterList {
			gotReward := arklib.ToArkFalt(sharedReward * vot.Weight)
			log.Infof("block reward: %v, my weigth: %v, my raword: %f", bot.Reward, vot.Weight, gotReward)
			var voter delegate.Vote
			db.Where("address = ?", vot.Address).First(&voter)
			mblock := delegate.VoteProfit{
				Block:   bot,
				Vote:    voter,
				Balance: voter.Balance,
				Reward:  gotReward,
				Payment: false,
			}
			db.Save(&mblock)
		}
	}
}

type SetPayoutOKInput struct {
	Stime  int64  `json:"start_time" form:"start_time" binding:"required"`
	Etime  int64  `json:"until_time" form:"until_time" binding:"required"`
	VoteId string `json:"vote_id" form:"vote_id" binding:"required"`
}

func SetPayoutOK(c *gin.Context) {
	inputs := SetPayoutOKInput{}
	if err := c.Bind(&inputs); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
		return
	}
	dhelp := delegate.VoteProfit{}
	dt := db.Begin()
	dt2 := dt.Table(dhelp.TableName()).Where("created_at >= ? AND created_at <= ? AND vote_id = ?", time.Unix(inputs.Stime, 0), time.Unix(inputs.Etime, 0), inputs.VoteId).Update("payment", true)
	if dt2.Error != nil {
		dt.Rollback()
		c.JSON(http.StatusBadGateway, gin.H{
			"error": dt2.Error,
		})
		return
	}
	dt.Commit()
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("row affected %d", dt2.RowsAffected),
	})
	return
}
