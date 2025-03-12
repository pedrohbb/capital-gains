package operations

import (
	. "github.com/pedrohbb/capital-gains/pkg"
)

type Operation struct {
	Operation string  `json:"operation"`
	UnitCost  Float64 `json:"unit-cost"`
	Quantity  int     `json:"quantity"`
}

type TaxOverOperation struct {
	Tax Float64 `json:"tax"`
}

type ChainOperations struct {
	Taxes       []TaxOverOperation `json:"taxes"`
	BuyWA       Float64            `json:"-"`
	Loss        Float64            `json:"-"`
	AccQuantity int                `json:"-"`
}