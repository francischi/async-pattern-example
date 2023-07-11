package pkg

import (
	"github.com/gin-gonic/gin"
	"golang/pkg/memberModule"
	"golang/pkg/orderModule"
)

func SetRouter(g *gin.Engine) {

	const baseGroup string = "/api"
	
	memberModule.SetRoute(g ,baseGroup)
	orderModule.SetRoute(g,baseGroup)
}