package ark

import (
	"time"
)

type ArkTransaction struct {
	ID        string `gorm:"type:varchar(100);primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	// TransactionHash string `gorm:"type:varchar(100);unique_index`
	TransactionTime time.Time
	FromAddr        string
	ToAddr          string
	BlockId         string
	Signature       string
	SenderPublicKey string
	Amount          float64
	Fee             int64
	SmartBridge     string
}
