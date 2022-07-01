package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rikimaru-tokyo/gin_sqlx_lesson02/model"
)

func In(c *gin.Context) {
	err := model.Insert()
	if err != nil {
		panic(err)
	}

	result, err := model.GetAll()

	if err != nil {
		panic(err)
	}
	c.JSON(200, result)
}

func Up(c *gin.Context) {
	err := model.Update()
	if err != nil {
		panic(err)
	}

	result, err := model.GetAll()

	if err != nil {
		panic(err)
	}
	c.JSON(200, result)
}

func Bulk(c *gin.Context) {
	err := model.BulkInsert()
	if err != nil {
		panic(err)
	}

	result, err := model.GetAll()

	if err != nil {
		panic(err)
	}
	c.JSON(200, result)
}
