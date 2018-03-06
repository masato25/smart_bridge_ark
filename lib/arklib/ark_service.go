package arklib

import (
	"fmt"
	"strconv"

	"github.com/masato25/ark-go/core"
	"github.com/masato25/smart_bridge_ark/app/model/delegate"
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
	conf := config.MyConfig().Ark
	if conf.DevNet != "prod" {
		core.EnvironmentParams.Network.Type = core.DEVNET
	} else {
		core.EnvironmentParams.Network.Type = core.MAINNET
	}
	arkclient = core.NewArkClientFromIP(conf.Host)
}

func GetVoters() core.DelegateVoters {
	// arkapi := core.NewArkClient(nil)
	voters, httpres, err := arkclient.GetDelegateVoters(genDelegateParams())
	log.Infof("voters: %v, http response: %v, error: %v", voters, httpres, err)
	return voters
}

type Account struct {
	Address   string  `json:"address"`
	PublicKey string  `json:"publicKey"`
	Balance   float64 `json:"balance"`
	Weight    float64 `json:"width"`
}

func CalculatorData(votes []delegate.Vote) []Account {
	var totallBalance float64
	var accts []struct {
		Username  string `json:"username"`
		Address   string `json:"address"`
		PublicKey string `json:"publicKey"`
		Balance   string `json:"balance"`
	}
	for _, acct := range votes {
		totallBalance += acct.Balance
		var atstruct struct {
			Username  string `json:"username"`
			Address   string `json:"address"`
			PublicKey string `json:"publicKey"`
			Balance   string `json:"balance"`
		}
		atstruct.Address = acct.Address
		atstruct.Balance = fmt.Sprintf("%v", acct.Balance)
		accts = append(accts, atstruct)
	}
	var acct core.DelegateVoters
	acct.Accounts = accts
	log.Infof("accounts: %v", acct)
	return CalculatorWeight(acct, totallBalance)
}

func CalculatorWeight(acct core.DelegateVoters, totallBalance float64) []Account {
	accounts := make([]Account, len(acct.Accounts))
	for indx, acct := range acct.Accounts {
		balance, _ := strconv.ParseFloat(acct.Balance, 64)
		Arkbalance := balance
		accounts[indx].Address = acct.Address
		accounts[indx].Balance = ToArkFalt(balance)
		accounts[indx].PublicKey = acct.PublicKey
		accounts[indx].Weight = float64((Arkbalance / totallBalance) * 100)
	}
	return accounts
}

func GetVotersAndWeight() []Account {
	voters := GetVoters()
	var totallBalance float64
	for _, acct := range voters.Accounts {
		balance, _ := strconv.ParseFloat(acct.Balance, 64)
		totallBalance += balance
	}
	accounts := CalculatorWeight(voters, totallBalance)
	return accounts
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
	tx := core.CreateTransaction(recepient, amount, "ARK-GOLang testing", content, "")
	payload := core.TransactionPayload{}
	payload.Transactions = append(payload.Transactions, tx)
	res, httpresponse, err := arkclient.PostTransaction(payload)
	if err != nil {
		return err
	}
	log.Debugf("res: %v, httpresponse: %v, err: %v", res, httpresponse, err)
	return
}
