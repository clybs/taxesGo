package main

import (
	"bufio"
	"fmt"
	"github.com/clybs/taxesGo/constants"
	"github.com/clybs/taxesGo/headers"
	"github.com/clybs/taxesGo/items"
	"os"
	"strconv"
)

var (
	hd         headers.Headers
	it         items.Items
	blankValue int
)

func calculateResults(v []items.Item) {
	var colTotal, colSalesTax float64

	for _, valRow := range v {
		// Calculate rows
		sQuantity := strconv.Itoa(valRow.Quantity)
		p := it.ComputeRowValue(valRow.Product, float64(valRow.Quantity)*valRow.Price)
		sPrice := strconv.FormatFloat(p, 'f', 2, 64)

		// Aggregate values
		colTotal += p
		colSalesTax += p - float64(valRow.Quantity)*valRow.Price

		var rowResult string

		// Display according to header sequence
		for i, valHeader := range hd.Values {
			switch valHeader {
			case constants.HeaderQuantity:
				rowResult += sQuantity
			case constants.HeaderProduct:
				rowResult += valRow.Product
			case constants.HeaderPrice:
				rowResult += sPrice
			}

			// Add separator
			if i < 2 {
				rowResult += ", "
			}
		}

		fmt.Println(rowResult)
	}
	fmt.Println("\nSales Taxes:", it.ComputeToFixed(colSalesTax, 2))
	fmt.Println("Total:", it.ComputeToFixed(colTotal, 2))
	resetData()
}

func main() {
	var headersValid bool

	scanner := bufio.NewScanner(os.Stdin)

	// Scan user input
	for scanner.Scan() {
		input := scanner.Text()

		// Check if input is set already
		if !headersValid {
			// Not yet set, check if headers are valid
			validHeader, entries := hd.IsValid(input)
			headersValid = validHeader

			// Output error if headers are not valid
			if !headersValid {
				fmt.Println(constants.ErrorHeadersNotValid)
			} else {
				// Save or reset the headers
				hd.SaveResetOrder(headersValid, entries)
			}
		} else {
			// Headers are already set
			// Capture the items section
			if len(input) == 0 {
				blankValue++
				if blankValue == 1 {
					// Calculate the data
					calculateResults(it.Rows)
					fmt.Println()

					// Reset the blank value count and item list
					resetData()
				}
			} else {
				_, err := it.Save(input, hd.Values)
				if err != nil {
					fmt.Println(err.Error())
				}
			}
		}
	}

	// Gotcha errors
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Reading standard input:", err)
	}
}

func resetData() {
	blankValue = 0
	it.ClearItemRow()
}
