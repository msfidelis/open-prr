package products

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"open-prr/models/products"
	"time"
)

type ResponseCreate struct {
	Status int `json:"status" binding:"required"`
}

type ResponseError struct {
	Message string `json:"message" binding:"required"`
	Status  int    `json:"status" binding:"required"`
}

func Create(c *gin.Context) {
	var response ResponseCreate
	var request products.Product

	// Request Validation
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, ResponseError{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	id, _ := uuid.NewUUID()

	creationTimestamp := time.Now()
	request.CreatedOn = creationTimestamp
	request.UpdatedOn = creationTimestamp

	request.Id = id.String()

	response.Status = http.StatusCreated
	c.JSON(http.StatusCreated, request)
}
