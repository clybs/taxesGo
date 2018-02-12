package items

import (
	"math"
)

// ComputeRowValue gets the row value
func (it *Items) ComputeRowValue(v string, amount float64) float64 {
	totalTax := it.computeTax10(v) + it.computeTax05(v)
	amountTotalTax := it.computeRoundSpecific(amount*totalTax, 0.05)
	totalAmount := amount + amountTotalTax

	return it.ComputeToFixed(totalAmount, 2)
}

// computeRound gets the rounded value
func (it *Items) computeRound(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

// computeRoundSpecific rounds to a specific value
func (it *Items) computeRoundSpecific(v, n float64) float64 {
	return float64(int64(v/n+0.5)) * n
}

// computeTax05 gets the 0.05 tax value
func (it *Items) computeTax05(v string) float64 {
	if it.isImported(v) {
		return 0.05
	}
	return 0
}

// computeTax10 gets the 0.1 tax value
func (it *Items) computeTax10(v string) float64 {
	if it.isBook(v) || it.isFood(v) || it.isMedicine(v) {
		return 0
	}
	return 0.10
}

// ComputeToFixed gets the fix computed value
func (it *Items) ComputeToFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(it.computeRound(num*output)) / output
}
