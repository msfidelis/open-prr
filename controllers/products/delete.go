package products

import (
	"net/http"
	"open-prr/models/products"
	"open-prr/pkg/orm"

	"open-prr/pkg/logger"

	"github.com/gin-gonic/gin"
)

type ResponseDelete struct {
	Status int `json:"status" binding:"required"`
}

func Delete(c *gin.Context) {
	var response ResponseDelete
	var product products.Product
	log := logger.Instance()

	id := c.Param("id")

	db, err := orm.GetInstance()

	if err != nil {
		panic(err)
	}

	result := db.Where("id = ?", id).Delete(&product)

	if result.Error != nil {
		log.Error().
			Str("id", id).
			Str("error", result.Error.Error()).
			Str("component", "products").
			Str("action", "delete").
			Str("status", "error").
			Msg("Product cannot be deleted")
	}

	log.Info().
		Str("id", id).
		Str("component", "products").
		Str("action", "delete").
		Str("status", "success").
		Msg("Product Deleted using Idempotency Method")

	c.JSON(http.StatusNoContent, response)
}
