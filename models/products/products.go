package products

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Id          string    `gorm:"column:id"`
	Name        string    `gorm:"column:name"`
	Description string    `gorm:"column:description"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}

func (Product) TableName() string {
	return "prr_products"
}
