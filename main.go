package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/masato25/smart_bridge_ark/app/controller"
	"github.com/masato25/smart_bridge_ark/app/setup"
	"github.com/masato25/smart_bridge_ark/app/views"
	"github.com/masato25/smart_bridge_ark/config"
	"github.com/masato25/smart_bridge_ark/lib/exchange_rate"
	log "github.com/sirupsen/logrus"
)

func main() {
	err := config.ReadConf()
	if err != nil {
		log.Error(err)
		os.Exit(3)
	}
	conf := config.MyConfig()
	exchange_rate.Init()
	setup.ConnDB()
	r := gin.Default()
	r.LoadHTMLGlob("app/views/**/*")
	r.Static("/assets", "./app/assets")
	r.HTMLRender = views.GetMultiRender()
	controller.Routes(r)
	r.Run(fmt.Sprintf(":%d", conf.Web.Port)) // listen and serve on 0.0.0.0:8080
}
