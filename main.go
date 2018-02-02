package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/masato25/smart_bridge_ark/app/controller"
	"github.com/masato25/smart_bridge_ark/app/setup"

	"github.com/masato25/smart_bridge_ark/app/views"
	"github.com/masato25/smart_bridge_ark/config"
	"github.com/masato25/smart_bridge_ark/lib/arklib"
	"github.com/masato25/smart_bridge_ark/lib/ethereumrpc"
	"github.com/masato25/smart_bridge_ark/lib/exchange_rate"
	log "github.com/sirupsen/logrus"
)

var myconf config.ViperConfig

func startService() {
	exchange_rate.Init()
	if myconf.Ark.Enable {
		log.Info("ark service will start")
		arklib.Conn()
	} else {
		log.Warn("ark service not enable")
	}
	if myconf.Ether.Enable {
		log.Info("ethereum service will start")
		ethereumrpc.Conn()
		ethereumrpc.SetTimeOut()
	} else {
		log.Warn("ethereum service not enable")
	}
	setup.ConnDB()
}

func main() {
	log.Info("[[[[[[[[[[[[[[[[[Smart Bridge]]]]]]]]]]]]]]]]]")
	err := config.ReadConf()
	if err != nil {
		log.Error(err)
		os.Exit(3)
	}
	myconf = config.MyConfig()
	startService()
	r := gin.Default()
	r.LoadHTMLGlob("app/views/**/*")
	r.Static("/assets", "./app/assets")
	r.HTMLRender = views.GetMultiRender()
	controller.Routes(r)
	r.Run(fmt.Sprintf(":%d", myconf.Web.Port)) // listen and serve on 0.0.0.0:8080
}
