package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rikimaru-tokyo/gin_sqlx_lesson03/model"
)

type TargetParameter struct {
	ID       int    `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Birthday string `json:"birthday" binding:"required"`
}

func GetTargetsAll(c *gin.Context) {

	result, err := model.FindTargetAll()
	if err != nil {
		c.AbortWithStatusJSON(500, err)
	}
	c.JSON(200, result)
}

func GetTargetSingle(c *gin.Context) {
	var id model.TargetID
	if err := c.ShouldBindUri(&id); err != nil {
		c.AbortWithStatusJSON(400, err)
	}

	result, err := model.FindTargetOne(id.ID)
	if err != nil {
		c.AbortWithStatusJSON(500, err)
	}
	c.JSON(200, result)
}

func PostSingleTarget(c *gin.Context) {
	var postBody model.TargetsTable

	if err := c.ShouldBindJSON(&postBody); err != nil {
		c.AbortWithStatusJSON(400, err)
	}

	result, err := model.InsertTargetOne(postBody)
	if err != nil {
		c.AbortWithStatusJSON(500, err)
	}

	type resp struct {
		result         string
		last_insert_id int
	}

	c.JSON(200, resp{result: "success", last_insert_id: result})
}

func PostBulkTargets(c *gin.Context) {
	var postBodies []model.TargetsTable

	if err := c.ShouldBindJSON(&postBodies); err != nil {
		c.AbortWithStatusJSON(400, err)
	}

	if err := model.InsertTargetBulk(postBodies); err != nil {
		c.AbortWithStatusJSON(500, err)

	}
	c.JSON(200, "Bulk insert success")
}

func PatchTarget(c *gin.Context) {
	var id model.TargetID
	var patchbody model.TargetUpdate

	if err := c.ShouldBindUri(&id); err != nil {
		c.AbortWithStatusJSON(400, err)
	}

	if err := c.ShouldBindJSON(&patchbody); err != nil {
		c.AbortWithStatusJSON(400, err)
	}

	if err := model.UpdateTarget(id.ID, patchbody); err != nil {
		c.AbortWithStatusJSON(500, err)
	}
	c.JSON(200, "update success")
}

func DeleteTarget(c *gin.Context) {
	var id int
	if err := c.ShouldBindUri(&id); err != nil {
		c.AbortWithStatusJSON(400, err)
	}

	if err := model.DeleteTarget(id); err != nil {
		c.AbortWithStatusJSON(500, err)
	}
	c.JSON(200, "delete success")
}
