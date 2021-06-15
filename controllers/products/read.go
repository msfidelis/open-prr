package products

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseDetail struct {
	Status int `json:"status" binding:"required"`
}

type ResponseList struct {
	Status int `json:"status" binding:"required"`
}

func Detail(c *gin.Context) {
	var response ResponseDetail
	response.Status = http.StatusOK
	c.JSON(http.StatusOK, response)
}

func List(c *gin.Context) {
	var response ResponseList
	response.Status = http.StatusOK
	c.JSON(http.StatusOK, response)
}
