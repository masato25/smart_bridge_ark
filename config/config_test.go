package config

import (
	"testing"

	log "github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSpec(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	Convey("Test Read Conf", t, func() {
		err := ReadConf()
		So(err, ShouldBeEmpty)
		conf := MyConfig()
		log.Infof("conf: %v", conf)
	})
}
