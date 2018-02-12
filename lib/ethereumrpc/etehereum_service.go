package ethereumrpc

import (
	"math/big"

	log "github.com/sirupsen/logrus"

	"context"
	"time"

	"github.com/ethereum/go-ethereum/rpc"
	"github.com/masato25/smart_bridge_ark/config"
	_ "github.com/onrik/ethrpc"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
)

var d = time.Now().Add(40000 * time.Millisecond)
var client *ethclient.Client

func checkConn() {
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()
	_, err := client.NetworkID(ctx)
	if err != nil {
		client = nil
	}
	Conn()
}

func SetTimeOut() {
	conf := config.MyConfig().Ether
	timeMs := conf.TimeOutMS
	d = time.Now().Add(time.Duration(timeMs) * time.Millisecond)
}

func getKey() (keystoredkey *keystore.Key, err error) {
	conf := config.MyConfig().Ether
	var keyjsonstring string
	keyjsonstring, err = conf.GetKeyString()
	if err != nil {
		log.Error(err)
		return keystoredkey, err
	}
	log.Debugf("getKey() keystring: %s", keyjsonstring)
	return keystore.DecryptKey([]byte(keyjsonstring), conf.KeyPassword)
}

func SendRawTransaction(sendToAddr string, amount float64) (signTx *types.Transaction, err error) {
	checkConn()
	conf := config.MyConfig().Ether
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()
	var unlockedKey *keystore.Key
	unlockedKey, err = getKey()
	if err != nil {
		log.Error(err)
		return
	}
	log.Debugf("%v %v %v", ctx, unlockedKey, err)
	sendToAddrComm := common.HexToAddress(sendToAddr)
	nonce, _ := client.NonceAt(ctx, unlockedKey.Address, nil)
	bdata := conf.DisplayMessage
	tx := types.NewTransaction(nonce, sendToAddrComm, big.NewInt(int64(amount*params.Ether)), big.NewInt(conf.GasFee), big.NewInt(10*params.Shannon), []byte(bdata))
	signTx, err = types.SignTx(tx, types.HomesteadSigner{}, unlockedKey.PrivateKey)
	if err != nil {
		log.Errorf("types.SignTx: %v", err)
		return
	}
	err = client.SendTransaction(ctx, signTx)
	if err != nil {
		log.Errorf("SendTransaction: %v", err)
	} else {
		log.Infof("signTx: %v", signTx)
	}
	return
}

func GetBalaceOf(address string) {
	checkConn()
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()
	addr := common.HexToAddress(address)
	blanceOf, err := client.BalanceAt(ctx, addr, nil)
	etherAmount := (float64(blanceOf.Int64()) / float64(params.Ether))
	log.Infof("ether blanceOf: %v, err: %v", etherAmount, err)
}

func Conn() error {
	conf := config.MyConfig().Ether
	rpcclient, err := rpc.DialHTTP(conf.RPCHost)
	if err != nil {
		log.Error(err)
		return err
	}
	client = ethclient.NewClient(rpcclient)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()
	netid, _ := client.NetworkID(ctx)
	log.Debugf("Network id: %d", netid)
	return nil
}
