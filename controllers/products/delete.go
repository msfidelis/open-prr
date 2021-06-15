package products

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseDelete struct {
	Status int `json:"status" binding:"required"`
}

func Delete(c *gin.Context) {
	var response ResponseDelete
	response.Status = http.StatusOK
	c.JSON(http.StatusOK, response)
}
