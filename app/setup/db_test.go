package setup

import (
	"testing"

	"github.com/masato25/smart_bridge_ark/config"
	log "github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSpec(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	Convey("Conn to db", t, func() {
		config.ReadConf("../../config")
		ConnDB(true)
	})
}
