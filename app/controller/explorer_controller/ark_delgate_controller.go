package explorer_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"github.com/masato25/ark-go/core"
	"github.com/masato25/smart_bridge_ark/app/model/blocks"
	"github.com/masato25/smart_bridge_ark/app/model/delegate"
	_ "github.com/masato25/smart_bridge_ark/app/model/ether"
	"github.com/masato25/smart_bridge_ark/config"
	"github.com/masato25/smart_bridge_ark/lib/arklib"
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
		db.Where("address = ?", Dvote.Address).FirstOrCreate(&Dvote)
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
		"conf": config.MyConfig().Ark,
	})
	return
}

func GetVoterController(c *gin.Context) {
	Dvote := []delegate.Vote{}
	db.Where("status = ?", true).Find(&Dvote)
	c.JSON(200, gin.H{
		"data": Dvote,
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

func GetBlockInfoByTimeLimit(c *gin.Context) {
	// fheight := arklib.GetCurrentHeight()
	fheight := 3720933
	var timestampLimit int64
	timestampLimit = 1520056906
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

func CalculatorEachBlocks(c *gin.Context) {
	blockResults := []blocks.Block{}
	db.Find(&blockResults)
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
