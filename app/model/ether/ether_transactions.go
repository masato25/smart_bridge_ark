package ether

import "time"

type EtherTransaction struct {
	ID        string `gorm:"type:varchar(100);primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	// TransactionHash string `gorm:"type:varchar(100);unique_index`
	FromAddr string
	ToAddr   string
	Amount   float64
	Data     string
}
