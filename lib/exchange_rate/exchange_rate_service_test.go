package exchange_rate

import (
	"testing"

	log "github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSpec(t *testing.T) {
	Init()
	log.SetLevel(log.DebugLevel)
	Convey("Get Rate from BTC to Ark", t, func() {
		rate, err := GetRate("ark", "ethereum")
		So(err, ShouldBeEmpty)
		log.Debug(rate)
		log.Debug(rate.Rate(1))
		So(rate.Rate(1), ShouldNotEqual, 0)
	})
}
