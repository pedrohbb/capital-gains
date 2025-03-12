package operations

import (
	. "github.com/pedrohbb/capital-gains/pkg"
)

type OperationService struct {
	taxFreeBucket Float64
	taxRate       Float64
}

func (svc *OperationService) ProcessOperationTaxes(opList []Operation) ChainOperations {
	stockResult := ChainOperations{
		Taxes:       []TaxOverOperation{},
		BuyWA:       0.00,
		Loss:        0.00,
		AccQuantity: 0,
	}

	for _, op := range opList {
		if op.Operation == "buy" {
			stockResult.BuyWA = svc.updateWeightedAverage(op, stockResult.BuyWA, stockResult.AccQuantity)
			stockResult.AccQuantity += op.Quantity
			stockResult.Taxes = append(stockResult.Taxes, TaxOverOperation{Tax: 0.00})

		} else if op.Operation == "sell" {
			tax := svc.calculateOperationTax(op, stockResult.BuyWA, stockResult.Loss)
			stockResult.Taxes = append(stockResult.Taxes, TaxOverOperation{Tax: tax})
			stockResult.AccQuantity -= op.Quantity

			// Tax and loss are mutually exclusive
			if tax == 0.00 {
				stockResult.Loss += (stockResult.BuyWA - op.UnitCost) * Float64(op.Quantity)
			}
		}
	}

	return stockResult
}

func (svc *OperationService) updateWeightedAverage(op Operation, currentAvg Float64, currentQuantity int) Float64 {
	return ((Float64(currentQuantity) * currentAvg) + (Float64(op.Quantity) * op.UnitCost)) / Float64(currentQuantity+op.Quantity)
}

func (svc *OperationService) calculateOperationTax(op Operation, currentAvg Float64, loss Float64) Float64 {
	profit := (op.UnitCost - currentAvg) * Float64(op.Quantity)
	netProfit := profit - loss

	totalValue := op.UnitCost * Float64(op.Quantity)
	if totalValue <= svc.taxFreeBucket {
		return Float64(0.00)
	}

	if netProfit > Float64(0.00) {
		return netProfit * svc.taxRate
	}

	return Float64(0.00)
}

func NewService(taxFreeBucket Float64, taxRate Float64) *OperationService {
	return &OperationService{taxFreeBucket: taxFreeBucket, taxRate: taxRate}
}
