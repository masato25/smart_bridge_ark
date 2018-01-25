package lib

import (
	"testing"

	"github.com/masato25/smart_bridge_ark/app/setup"
	"github.com/masato25/smart_bridge_ark/config"
	"github.com/masato25/smart_bridge_ark/lib/arklib"
	log "github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSpec(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	config.ReadConf("../../config")
	setup.ConnDB()
	Convey("Get Voters", t, func() {
		arklib.Conn()
		arklib.GetVoters()
	})
}
