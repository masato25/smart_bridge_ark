package database

import (
	"testing"

	"github.com/masato25/smart_bridge_ark/app/model/ark"
	"github.com/masato25/smart_bridge_ark/app/setup"
	"github.com/masato25/smart_bridge_ark/config"
	log "github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
)

func TestArkDB(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	config.ReadConf("../../config")
	setup.ConnDB()
	db := setup.GetConn()
	Convey("test insert ark record", t, func() {
		at := ark.ArkTransaction{
			FromAddr:        "D00000000000000000000000000000000s",
			ToAddr:          "D00000000000000000000000000000000s",
			TransactionHash: "000000000000000000000000000000000000000000000000000000000000000000",
			Amount:          0.1,
		}
		db.Create(&at)
		var at2 []ark.ArkTransaction
		db.Find(&at2)
		log.Infof("arks: %v", at2)
	})
}
