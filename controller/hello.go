package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rikimaru-tokyo/gin_sqlx_lesson02/model"
)

func Hello(c *gin.Context) {
	c.JSON(200, gin.H{"Hello": "world"})
}

func GetMembers(c *gin.Context) {
	result, err := model.UseSelect(2)
	if err != nil {
		panic(err)
	}
	c.JSON(200, result)
}

func GetMemberExists(c *gin.Context) {
	member_id, _ := strconv.Atoi(c.Param("id"))

	result, err := model.UseGet(member_id)
	if err != nil {
		panic(err)
	}
	c.JSON(200, result)
}

func PostRanks(c *gin.Context) {
	result, err := model.InsertRank()
	if err != nil {
		panic(err)
	}
	c.JSON(200, result)
}

func DeleteRanks(c *gin.Context) {
	result, err := model.DeleteRank()
	if err != nil {
		panic(err)
	}
	c.JSON(200, result)
}
