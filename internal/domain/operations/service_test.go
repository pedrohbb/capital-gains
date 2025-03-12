package operations

import (
	"testing"
)

func TestProcessOperationTaxes(t *testing.T) {
	svc := NewService(20000.00, 0.20)

	testCases := []struct {
		input  []Operation
		output ChainOperations
	}{
		{
			input: []Operation{
				{
					Operation: "buy", UnitCost: 10.00, Quantity: 100,
				},
				{
					Operation: "sell", UnitCost: 15.00, Quantity: 50,
				},
				{
					Operation: "sell", UnitCost: 15.00, Quantity: 50,
				},
			},
			output: ChainOperations{Taxes: []TaxOverOperation{{Tax: 0.00}, {Tax: 0.00}, {Tax: 0.00}}},
		},
		{
			input: []Operation{
				{
					Operation: "buy", UnitCost: 10.00, Quantity: 10000,
				},
				{
					Operation: "sell", UnitCost: 20.00, Quantity: 5000,
				},
				{
					Operation: "sell", UnitCost: 5.00, Quantity: 5000,
				},
			},
			output: ChainOperations{Taxes: []TaxOverOperation{{Tax: 0.00}, {Tax: 10000.00}, {Tax: 0.00}}},
		},
		{
			input: []Operation{
				{
					Operation: "buy", UnitCost: 10.00, Quantity: 10000,
				},
				{
					Operation: "buy", UnitCost: 25.00, Quantity: 5000,
				},
				{
					Operation: "sell", UnitCost: 15.00, Quantity: 10000,
				},
				{
					Operation: "sell", UnitCost: 25.00, Quantity: 5000,
				},
			},
			output: ChainOperations{Taxes: []TaxOverOperation{{Tax: 0.00}, {Tax: 0.00}, {Tax: 0.00}, {Tax: 10000.00}}},
		},
		{
			input: []Operation{
				{
					Operation: "buy", UnitCost: 10.00, Quantity: 10000,
				},
				{
					Operation: "sell", UnitCost: 50.00, Quantity: 10000,
				},
				{
					Operation: "buy", UnitCost: 20.00, Quantity: 10000,
				},
				{
					Operation: "sell", UnitCost: 50.00, Quantity: 10000,
				},
			},
			output: ChainOperations{Taxes: []TaxOverOperation{{Tax: 0.00}, {Tax: 80000.00}, {Tax: 0.00}, {Tax: 60000.00}}},
		},
	}

	for _, tc := range testCases {
		t.Run("Testing ProcessOperationTaxes", func(t *testing.T) {
			got := svc.ProcessOperationTaxes(tc.input)
			for i, r := range got.Taxes {

				if r.Tax != tc.output.Taxes[i].Tax {
					t.Errorf("For input %v, expected %v but got %v", tc.input, tc.output, got)
				}
			}
		})
	}
}
