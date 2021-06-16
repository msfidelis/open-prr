package products

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"open-prr/models/products"
	"open-prr/pkg/logger"
	"open-prr/pkg/orm"
)

type RequestCreate struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type ResponseCreate struct {
	Status      int       `json:"status" binding:"required"`
	Id          string    `json:"id" binding:"required"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdateAt    time.Time `json:"updated_at"`
}

func Create(c *gin.Context) {
	var response ResponseCreate
	var request RequestCreate

	log := logger.Instance()

	// Request Validation
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, ResponseError{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	db, err := orm.GetInstance()

	if err != nil {
		panic(err)
	}

	// Check Product Name
	var productNameCheck products.Product
	resultCheck := db.Where("name = ?", request.Name).First(&productNameCheck)

	if resultCheck.RowsAffected != 0 || productNameCheck.Id != "" {
		log.Warn().
			Str("name", request.Name).
			Str("description", request.Description).
			Str("error", "product already exists").
			Str("component", "products").
			Str("action", "creation").
			Str("status", "error").
			Msg("Error to create new product")

		c.JSON(http.StatusBadRequest, &ResponseError{
			Status:  http.StatusBadRequest,
			Message: fmt.Sprintf("Product %s already exists", request.Name),
		})
		return
	}

	fmt.Println("Produto do banco", productNameCheck.Id)
	fmt.Println("Quantidade", resultCheck.RowsAffected)

	// New Product
	var product products.Product

	id, _ := uuid.NewUUID()
	product.Id = id.String()
	product.Name = request.Name
	product.Description = request.Description

	// Insert Action
	result := db.Create(&product)

	if result.Error != nil {
		log.Error().
			Str("name", product.Name).
			Str("description", product.Description).
			Str("error", result.Error.Error()).
			Str("component", "products").
			Str("action", "creation").
			Str("status", "error").
			Msg("Error to create new product")

		c.JSON(http.StatusInternalServerError, &ResponseError{
			Status:  http.StatusInternalServerError,
			Message: "Try again later",
		})
		return
	}

	log.Info().
		Str("id", product.Id).
		Str("name", product.Name).
		Str("description", product.Description).
		Str("component", "products").
		Str("action", "creation").
		Str("status", "success").
		Msg("New product created")

	response.Status = http.StatusCreated
	response.Id = product.Id
	response.Name = product.Name
	response.Description = product.Description
	response.CreatedAt = product.CreatedAt
	response.UpdateAt = product.UpdatedAt

	c.JSON(http.StatusCreated, response)
}
