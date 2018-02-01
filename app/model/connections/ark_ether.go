package connections

import (
	"github.com/jinzhu/gorm"
	"github.com/masato25/smart_bridge_ark/app/model/ark"
	"github.com/masato25/smart_bridge_ark/app/model/ether"
)

type ArkEther struct {
	gorm.Model
	ArkT               ark.ArkTransaction
	EtherT             ether.EtherTransaction
	ArkTransactionID   string
	EtherTransactionID string
}
