package model

import (
	"database/sql"

	"github.com/shopspring/decimal"
)

type Product struct {
	Price       decimal.Decimal `json:"price"`
	Description sql.NullString  `json:"description"`
}
