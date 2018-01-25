package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/masato25/smart_bridge_ark/app/controller/explorer_controller"
	"github.com/masato25/smart_bridge_ark/app/controller/hello_controller"
)

func Routes(route *gin.Engine) {
	hello_controller.Route(route)
	explorer_controller.Route(route)
}
