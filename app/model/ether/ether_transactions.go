package ether

import "github.com/jinzhu/gorm"

type EtherTransaction struct {
	gorm.Model
	FromAddr        string
	ToAddr          string
	TransactionHash string
	Amount          float64
	Data            string
}
