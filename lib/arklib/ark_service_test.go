package arklib

import (
	"testing"

	log "github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSpec(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	Convey("Get Voters", t, func() {
		Conn()
		GetVoters()
	})
}
