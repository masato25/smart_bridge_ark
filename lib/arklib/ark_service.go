package arklib

import (
	"github.com/masato25/ark-go/core"
	"github.com/masato25/smart_bridge_ark/config"
	log "github.com/sirupsen/logrus"
)

var arkclient *core.ArkClient

func genDelegateParams() core.DelegateQueryParams {
	conf := config.MyConfig().Ark
	publickey := conf.PublicKey
	username := conf.UserName
	return core.DelegateQueryParams{
		UserName:  username,
		PublicKey: publickey,
	}
}

// https://blog.ark.io/devnet-node-setup-configuration-15f06328e8b7
func Conn() {
	arkclient = core.NewArkClient(nil).SetActiveConfiguration(core.DEVNET)
}

func GetVoters() {
	// arkapi := core.NewArkClient(nil)
	voters, _, _ := arkclient.GetDelegateVoters(genDelegateParams())
	log.Infof("voters: %v", voters)
}

func TransferTo(recepient string, amount int64) (err error) {
	conf := config.MyConfig().Ark
	var content string
	content, err = conf.GetPassString()
	if err != nil {
		log.Error(err)
		return
	}
	//create and send tx
	tx := core.CreateTransaction(recepient, 1, "ARK-GOLang testing", content, "")
	payload := core.TransactionPayload{}
	payload.Transactions = append(payload.Transactions, tx)
	res, httpresponse, err := arkclient.PostTransaction(payload)
	if err != nil {
		return err
	}
	log.Debugf("res: %v, httpresponse: %v, err: %v", res, httpresponse, err)
	return
}
