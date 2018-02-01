package arklib

import (
	"testing"

	log "github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSpec(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	Convey("Conver big number to Ark", t, func() {
		result := ToArk(int64(100000000))
		log.Debug(result)
	})
}
