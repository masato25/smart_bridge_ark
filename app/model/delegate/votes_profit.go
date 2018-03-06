package delegate

import (
	"github.com/jinzhu/gorm"
	"github.com/masato25/smart_bridge_ark/app/model/blocks"
)

type VoteProfit struct {
	gorm.Model
	Block   blocks.Block `json:"-"`
	Vote    Vote         `json:"-"`
	Balance float64
	Reward  float64
	Payment bool
	VoteID  string `gorm:"unique_index:uniq_profit_indx"`
	BlockID string `gorm:"unique_index:uniq_profit_indx"`
}

func (self *VoteProfit) TableName() string {
	return "vote_profits"
}
