package items

import (
	"testing"
)

func TestComputeRowValue(t *testing.T) {
	var it Items
	tables := []struct {
		val    string
		amount float64
		result float64
	}{
		{"", 1.0, 1.1},
		{"chocolate slice", 1.0, 1.0},
		{"imported birthday cake", 1.0, 1.05},
		{"candies and cookies", 1.0, 1.0},
		{"imported bottle of wine", 1.0, 1.05},
		{"imported box of shoes", 1.0, 1.15},
	}

	for _, table := range tables {
		result := it.ComputeRowValue(table.val, table.amount)
		if result != table.result {
			t.Errorf("ComputeRowValue(%v, %v) was incorrect, got: %v, want: %v.", table.val, table.amount, result, table.result)
		}
	}
}

func TestComputeTax05(t *testing.T) {
	var it Items
	tables := []struct {
		val    string
		result float64
	}{
		{"", 0},
		{"chocolate slice", 0},
		{"imported birthday cake", 0.05},
		{"candies and cookies", 0},
		{"imported bottle of wine", 0.05},
		{"imported box of shoes", 0.05},
	}

	for _, table := range tables {
		result := it.computeTax05(table.val)
		if result != table.result {
			t.Errorf("computeTax05(%v) was incorrect, got: %v, want: %v.", table.val, result, table.result)
		}
	}
}

func TestComputeTax10(t *testing.T) {
	var it Items
	tables := []struct {
		val    string
		amount float64
		result float64
	}{
		{"", 1.0, 0.1},
		{"chocolate slice", 1.0, 0},
		{"birthday cake", 1.0, 0},
		{"candies and cookies", 1.0, 0},
		{"bottle of wine", 1.0, 0},
		{"box of shoes", 1.0, 0.1},
	}

	for _, table := range tables {
		result := it.computeTax10(table.val)
		if result != table.result {
			t.Errorf("computeTax10(%v) was incorrect, got: %v, want: %v.", table.val, result, table.result)
		}
	}
}
