package ark

import (
	"testing"

	"github.com/masato25/smart_bridge_ark/app/setup"
	"github.com/masato25/smart_bridge_ark/config"
	"github.com/masato25/smart_bridge_ark/lib/arklib"
	log "github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
)

func TestArkService(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	config.ReadConf("../../../config")
	setup.ConnDB()
	Convey("Get Voters", t, func() {
		arklib.Conn()
		arklib.GetVoters()
		aktransaction := arklib.GetTransactionsReceive(config.MyConfig().Ark.Address)
		if aktransaction != nil {
			for _, tx := range *aktransaction {
				log.Debugf("resp tx: %v", tx.ToJSON())
			}
		} else {
			log.Debug("transaction size is 0")
		}
	})
}

// func TestArkTransactionDB(t *testing.T) {
// 	log.SetLevel(log.DebugLevel)
// 	config.ReadConf("../../../config")
// 	setup.ConnDB()
// 	db := setup
// 	Convey("Get Voters", t, func() {
// 		arklib.Conn()
// 		arklib.GetVoters()
// 		aktransaction := arklib.GetTransactions()
// 		for _, tx := range aktransaction {
// 			log.Debugf("resp tx: %v", tx.ToJSON())
// 		}
// 	})
// }
