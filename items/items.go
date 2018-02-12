package items

import (
	"errors"
	"github.com/clybs/taxesGo/constants"
	"github.com/clybs/taxesGo/utils"
	"strconv"
)

var ut utils.Utils

// Items struct containing array of rows
type Items struct {
	Rows []Item
}

// Item struct for the row
type Item struct {
	Quantity int
	Product  string
	Price    float64
}

// ClearItemRow clears an entire row value
func (it *Items) ClearItemRow() {
	item := make([]Item, 0)
	it.Rows = item
}

// IsValid checks if row entries are valid
func (it *Items) IsValid(v string, h []string) (bool, Item, error) {
	// Create container for entries
	var row Item

	// Get the entries
	entries := ut.GetEntries(v)

	// Check if length is 3
	if len(entries) != constants.ColumnCount {
		return false, row, errors.New(constants.ErrorWrongNumberOfArguments)
	}

	// Extract data
	for i, val := range h {
		switch val {
		case constants.HeaderQuantity:
			// Process quantity types
			iv, err := strconv.Atoi(entries[i])
			if err != nil {
				return false, row, errors.New(constants.ErrorItemQuantityNotValid)
			}
			row.Quantity = iv
		case constants.HeaderPrice:
			// Process price types
			iv, err := strconv.ParseFloat(entries[i], 64)
			if err != nil {
				return false, row, errors.New(constants.ErrorItemPriceNotValid)
			}
			row.Price = iv
		default:
			// Do not process at all
			row.Product = entries[i]
		}
	}

	return true, row, nil
}

// Save keeps row entry in a stack
func (it *Items) Save(v string, h []string) ([]Item, error) {
	// Check first if entry is valid
	valid, row, err := it.IsValid(v, h)
	if err != nil {
		return it.Rows, err
	}

	if valid {
		// Save the row
		it.Rows = append(it.Rows, row)
	}

	return it.Rows, nil
}
