package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rikimaru-tokyo/gin_sqlx_lesson02/controller"
)

func LoadRouter(engine *gin.Engine) {
	v1 := engine.Group("/v1")
	{
		v1.GET("/health_check", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "OK"})
		})
		v1.GET("/members", controller.GetMembers)
		v1.GET("/member/exists/:id", controller.GetMemberExists)
		v1.POST("/members", controller.PostRanks)
		v1.DELETE("/members", controller.DeleteRanks)
	}

	hoge := v1.Group("/hoge")
	{
		hoge.POST("/in", controller.In)
		hoge.PATCH("/up", controller.Up)
		hoge.POST("/bulk", controller.Bulk)
	}
}
