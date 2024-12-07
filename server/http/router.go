package http

import (
	"github.com/gin-gonic/gin"
	httpServer "github.com/lascape/gopkg/server/http"
	_ "github.com/lascape/goweb/docs" // 这里引入生成的 docs 包
	"github.com/lascape/goweb/server/http/api"
	files "github.com/swaggo/files"   // swagger embed files
	_ "github.com/swaggo/gin-swagger" // gin-swagger middleware
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	httpServer.AddRouter(func(engine *gin.Engine) {
		engine.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler))
		v1 := engine.Group("/api/v1")
		{
			goWeb := api.NewGoWebController()
			v1.GET("/go/web", goWeb.GoWeb)
		}
	})
}
