package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rikimaru-tokyo/gin_sqlx_lesson03/router"
)

func main() {
	engine := gin.Default()
	router.LoadRouter(engine)
	if err := engine.Run(":8080"); err != nil {
		panic(err)
	}
}
