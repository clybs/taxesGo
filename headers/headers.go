package headers

import (
	"github.com/clybs/taxesGo/constants"
	"github.com/clybs/taxesGo/utils"
)

var ut utils.Utils

// Headers struct to use
type Headers struct {
	Values []string
}

// isStringInSlice checks if a string is in slice
func (hd *Headers) isStringInSlice(v string, list []string) bool {
	for _, val := range list {
		if val == v {
			return true
		}
	}
	return false
}

// IsValid checks if header values are valid
func (hd *Headers) IsValid(v string) (bool, []string) {
	// Get the entries
	entries := ut.GetEntries(v)

	// Check if length is 3
	if len(entries) != constants.ColumnCount {
		return false, entries
	}

	// Check if expected headers are found
	priceFound := hd.isStringInSlice(constants.HeaderPrice, entries)
	productFound := hd.isStringInSlice(constants.HeaderProduct, entries)
	quantityFound := hd.isStringInSlice(constants.HeaderQuantity, entries)

	headersFound := priceFound && productFound && quantityFound

	return headersFound, entries
}

// SaveResetOrder decides if operation is to save or reset header values
func (hd *Headers) SaveResetOrder(Save bool, headers []string) {
	// Reset header value
	hd.Values = make([]string, 0)

	// Save value if found
	if Save {
		hd.Values = headers
	}
}
