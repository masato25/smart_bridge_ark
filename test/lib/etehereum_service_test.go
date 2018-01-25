package lib

import (
	"testing"

	"github.com/masato25/smart_bridge_ark/app/setup"
	"github.com/masato25/smart_bridge_ark/config"
	"github.com/masato25/smart_bridge_ark/lib/ethereumrpc"
	log "github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
)

func TestEtehereumService(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	config.ReadConf("../../config")
	setup.ConnDB()
	Convey("Test Ethereum net", t, func() {
		ethereumrpc.Conn()
		ethereumrpc.GetBalaceOf("0x955f26cde0a01b7182333a922d5f6b82aad516ed")
		ethereumrpc.SendRawTransaction("0x483e8C48b216687818dB732aF22874d722db68B7", 0.0000001)
	})
}
