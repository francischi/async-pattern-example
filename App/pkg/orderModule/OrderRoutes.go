package orderModule

import(
	"github.com/gin-gonic/gin"
	mw "golang/pkg/middleWare"
)

func SetRoute(g *gin.Engine , baseGroup string){

	orderGroup := g.Group(baseGroup+"/order")
	orderGroup.POST("",
		func(ctx *gin.Context){mw.InitJwtMiddleWare().ConfirmToken(ctx)}, 
		func(ctx *gin.Context){InitialOrderController().Add(ctx)},
	)

}