package connpassrouter

import (
	"github.com/gin-gonic/gin"
	middleware "github.com/pokotyan/connpass-map-api/middleware/logging"
	"github.com/pokotyan/connpass-map-api/routes/connpass/env"
	"github.com/pokotyan/connpass-map-api/routes/connpass/event"
)

func Init(router *gin.Engine) {
	connpass := router.Group("/connpass")
	{
		connpass.POST("/event", event.GetEvent, middleware.Logging())
		connpass.POST("/slack", event.PostSlack, middleware.Logging())
		connpass.POST("/env", env.Set, middleware.Logging())
		connpass.POST("/env/dialog", env.OpenDialog, middleware.Logging())
	}
}
