package example

import (
	"log"
	"math/big"

	"github.com/onrik/ethrpc"
)

func run() {
	client := ethrpc.NewEthRPC("http://0.0.0.0:8545")
	v, err := client.Web3ClientVersion()
	if err != nil {
		log.Fatal(err)
	}

	log.Print(v)
	// Send 1 eth
	txid, err := client.EthSendTransaction(ethrpc.T{
		From:  "0x6247cf0412c6462da2a51d05139e2a3c6c630f0a",
		To:    "0xcfa202c4268749fbb5136f2b68f7402984ed444b",
		Value: big.NewInt(1000000000000000000),
	})
	log.Print(txid)
}
