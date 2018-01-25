package ark

import "github.com/jinzhu/gorm"

type ArkTransaction struct {
	gorm.Model
	FromAddr        string
	ToAddr          string
	TransactionHash string
	Amount          float64
	SmartBridge     string
}
