package model

import (
	"database/sql"
	"time"

	"github.com/shopspring/decimal"
)

type Product struct {
	ID          int             `json:"id"`
	CategoryID  sql.NullInt32   `json:"categoryId"`
	Name        string          `json:"name"`
	Price       decimal.Decimal `json:"price"`
	Description sql.NullString  `json:"description"`
	Stock       int             `json:"stock"`
	CreatedAt   time.Time       `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt   time.Time       `json:"updatedAt" gorm:"autoUpdateTime"`
}

type ProductDto struct {
	ID          int       `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Price       int       `json:"price" binding:"min=10000"`
	Description *string   `json:"description"`
	Stock       int       `json:"stock" binding:"min=1"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func (p *ProductDto) FillFromModel(model Product) {
	p.ID = model.ID
	p.Name = model.Name
	p.Price = int(model.Price.IntPart())
	if model.Description.Valid {
		p.Description = &model.Description.String
	}
	p.Stock = model.Stock
	p.CreatedAt = model.CreatedAt
	p.UpdatedAt = model.UpdatedAt
}

func (p ProductDto) ToModel() Product {
	model := Product{
		ID:        p.ID,
		Name:      p.Name,
		Price:     decimal.NewFromInt(int64(p.Price)),
		Stock:     p.Stock,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
	if p.Description != nil {
		model.Description.String = *p.Description
		model.Description.Valid = true
	}

	return model
}
