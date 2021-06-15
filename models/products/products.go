package products

import (
	"gorm.io/gorm"
	"time"
)

type Product struct {
	gorm.Model
	Id          string    `json:"id" gorm:"column:id primaryKey"`
	Name        string    `json:"name" binding:"required" gorm:"column:name"`
	Description string    `json:"description" gorm:"column:description"`
	CreatedOn   time.Time `json:"created_on" gorm:"column:created_on"`
	UpdatedOn   time.Time `json:"updated_on" gorm:"column:updated_on"`
}

func (Product) TableName() string {
	return "prr_products"
}
