package products

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseUpdate struct {
	Status int `json:"status" binding:"required"`
}

func Update(c *gin.Context) {
	var response ResponseUpdate
	response.Status = http.StatusOK
	c.JSON(http.StatusOK, response)
}
