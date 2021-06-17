package products

import (
	"net/http"
	"open-prr/models/products"
	"open-prr/pkg/orm"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type ResponseDelete struct {
	Status int `json:"status" binding:"required"`
}

func Delete(c *gin.Context) {
	var response ResponseDelete
	var product products.Product

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
		Msg("Product Deleted")

	// log := logger.Instance()
	c.JSON(http.StatusNoContent, response)
}
