package products

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"open-prr/models/products"
	"open-prr/pkg/logger"
	"open-prr/pkg/orm"
)

type ResponseDetail struct {
	Id          string    `json:"id" binding:"required"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdateAt    time.Time `json:"updated_at"`
}

type ResponseList struct {
	Status   int              `json:"status" binding:"required"`
	Results  int64            `json:"results" binding:"required"`
	Products []ResponseDetail `json:"products" binding:"required"`
}

func List(c *gin.Context) {
	var response ResponseList
	var productsList []products.Product

	log := logger.Instance()

	db, err := orm.GetInstance()

	if err != nil {
		panic(err)
	}

	result := db.Find(&productsList)

	if result.Error != nil {
		log.Warn().
			Str("error", result.Error.Error()).
			Str("component", "products").
			Str("action", "list").
			Str("status", "error").
			Msg("Product list cannot be not retrieve")

		c.JSON(http.StatusNotFound, ResponseError{
			Status:  http.StatusNotFound,
			Message: fmt.Sprintf("Products not found"),
		})
	}

	for _, product := range productsList {
		response.Products = append(response.Products, ResponseDetail{
			Id:          product.Id,
			Name:        product.Name,
			Description: product.Description,
			CreatedAt:   product.CreatedAt,
			UpdateAt:    product.UpdatedAt,
		})
	}

	response.Status = http.StatusOK
	response.Results = result.RowsAffected
	c.JSON(http.StatusOK, response)
}

func Detail(c *gin.Context) {
	var response ResponseDetail
	var product products.Product

	log := logger.Instance()

	id := c.Param("id")

	db, err := orm.GetInstance()

	if err != nil {
		panic(err)
	}

	result := db.Where("id = ?", id).First(&product)

	if result.Error != nil {
		log.Warn().
			Str("id", id).
			Str("error", result.Error.Error()).
			Str("component", "products").
			Str("action", "list").
			Str("status", "error").
			Msg("Product cannot be not found")

		c.JSON(http.StatusNotFound, ResponseError{
			Status:  http.StatusNotFound,
			Message: fmt.Sprintf("Product %s not found", id),
		})

		return
	}

	response.Id = product.Id
	response.Name = product.Name
	response.Description = product.Description
	response.CreatedAt = product.CreatedAt
	response.UpdateAt = product.UpdatedAt

	c.JSON(http.StatusOK, response)
}
