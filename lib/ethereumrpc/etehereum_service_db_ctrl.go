package ethereumrpc

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/jinzhu/gorm"
	"github.com/masato25/smart_bridge_ark/app/model/ark"
	"github.com/masato25/smart_bridge_ark/app/model/connections"
	"github.com/masato25/smart_bridge_ark/app/model/ether"
	"github.com/masato25/smart_bridge_ark/app/setup"
	"github.com/masato25/smart_bridge_ark/config"
	_ "github.com/onrik/ethrpc"
	log "github.com/sirupsen/logrus"
)

var db *gorm.DB

func SyncSendEth(tx ark.ArkTransaction) (err error) {
	db = setup.GetConn()
	checkArkSend(tx.ID)
	var ak2 ark.ArkTransaction
	ak2, err = checkArkSend(tx.ID)
	if err != nil {
		log.Errorf("checkArkSend error: %s", err.Error())
		return
	}
	var addr string
	addr, err = checkEthInfo(ak2)
	if err != nil {
		log.Errorf("checkEthInfo error: %s", err.Error())
		// will do noting for all parse errors
		return nil
	}
	var send bool
	send, err = checkEthSend(ak2, addr)
	if err != nil {
		log.Errorf("checkEthSend error: %s", err.Error())
		return
	}
	if send {
		var ethrecord ether.EtherTransaction
		ethrecord, err = sendEthToUser(tx)
		if err != nil {
			log.Errorf("sendEthToUser error: %s", err.Error())
			return
		}
		newarketh := connections.ArkEther{
			EtherTransactionID: ethrecord.ID,
			ArkTransactionID:   tx.ID,
		}
		dt := db.Save(&newarketh)
		if dt.Error != nil {
			err = dt.Error
		}
	}
	return
}

func checkArkSend(txid string) (tx ark.ArkTransaction, err error) {
	arkt := ark.ArkTransaction{
		ID: txid,
	}
	arks := []ark.ArkTransaction{}
	db.Where(&arkt).Find(&arks)
	if len(arks) == 0 {
		err = fmt.Errorf("Transction id not found, please contact administrator.")
		return
	}
	tx = arks[0]
	return
}

type SmartBridgeEth struct {
	Address string `json:"eth"`
}

func checkEthInfo(tx ark.ArkTransaction) (addr string, err error) {
	dstruct := SmartBridgeEth{}
	err = json.Unmarshal([]byte(tx.SmartBridge), &dstruct)
	if err != nil {
		return
	}
	addr = dstruct.Address
	return
}

func checkEthSend(tx ark.ArkTransaction, ethaddr string) (status bool, err error) {
	arks := []connections.ArkEther{}
	db.Where("ark_transaction_id = ?", tx.ID).Find(&arks)
	if len(arks) == 0 {
		status = true
	} else {
		status = false
	}
	return
}

func sendEthToUser(tx ark.ArkTransaction) (ethrecord ether.EtherTransaction, err error) {
	var addr string
	addr, err = checkEthInfo(tx)
	if err != nil {
		return
	}
	var ethtx *types.Transaction
	amount := 0.0001
	ethtx, err = SendRawTransaction(addr, amount)
	if err != nil {
		return
	}
	rawtx := *ethtx
	rawtx.String()
	ethrecord = ether.EtherTransaction{
		ID:       fmt.Sprintf("0x%x", rawtx.Hash()),
		FromAddr: config.MyConfig().Ether.Address,
		ToAddr:   addr,
		Amount:   amount,
		Data:     string(rawtx.Data()),
	}
	dt := db.Create(&ethrecord)
	if dt.Error != nil {
		err = dt.Error
		return
	}
	return
}
