package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rikimaru-tokyo/gin_sqlx_lesson03/controller"
)

func LoadRouter(engine *gin.Engine) {
	v1 := engine.Group("/v1")
	{
		v1.GET("/health_check", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "OK"})
		})

		tgt := v1.Group("/target")
		{
			tgt.GET("", controller.GetTargetsAll)
			tgt.POST("", controller.PostSingleTarget)
			tgt.POST("/bulk", controller.PostBulkTargets)
			tgt.PATCH(":id", controller.PatchTarget)
			tgt.DELETE(":id", controller.DeleteTarget)
		}
	}
}
