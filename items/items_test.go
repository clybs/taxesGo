package items

import (
	"testing"

	"github.com/clybs/taxesGo/constants"
)

func TestClearItemRow(t *testing.T) {
	tables := []struct {
		val     string
		headers []string
		result  bool
	}{
		{"", []string{}, false},
		{"1, book, 1.5", []string{constants.HeaderQuantity, constants.HeaderProduct, constants.HeaderPrice}, true},
		{"1, book, 1.5", []string{constants.HeaderQuantity, constants.HeaderPrice, constants.HeaderProduct}, false},
		{"1, 1.5, book", []string{constants.HeaderQuantity, constants.HeaderProduct, constants.HeaderPrice}, false},
		{"1.2, book, 1.5", []string{constants.HeaderQuantity, constants.HeaderProduct, constants.HeaderPrice}, false},
		{"1.5, 2, book", []string{constants.HeaderPrice, constants.HeaderQuantity, constants.HeaderProduct}, true},
	}

	for _, table := range tables {
		var it Items
		_, _ = it.Save(table.val, table.headers)
		it.ClearItemRow()
		if len(it.Rows) > 0 {
			t.Errorf("ClearItemRow() was incorrect, got: %v, want: %v.", len(it.Rows), 0)
		}

	}
}

func TestIsValidItem(t *testing.T) {
	var it Items
	tables := []struct {
		val     string
		headers []string
		result  bool
	}{
		{"", []string{}, false},
		{"1, book, 1.5", []string{constants.HeaderQuantity, constants.HeaderProduct, constants.HeaderPrice}, true},
		{"1, book, 1.5", []string{constants.HeaderQuantity, constants.HeaderPrice, constants.HeaderProduct}, false},
		{"1, 1.5, book", []string{constants.HeaderQuantity, constants.HeaderProduct, constants.HeaderPrice}, false},
		{"1.2, book, 1.5", []string{constants.HeaderQuantity, constants.HeaderProduct, constants.HeaderPrice}, false},
		{"1.5, 2, book", []string{constants.HeaderPrice, constants.HeaderQuantity, constants.HeaderProduct}, true},
	}

	for _, table := range tables {
		result, _, err := it.IsValid(table.val, table.headers)
		if err != nil && table.result == true {
			t.Errorf("IsValid(%v, %v) was incorrect, got: %v, want: %v.", table.val, table.headers, err, nil)
			break
		}
		if result != table.result {
			t.Errorf("IsValid(%v, %v) was incorrect, got: %v, want: %v.", table.val, table.headers, result, table.result)
		}
	}
}

func TestSaveItem(t *testing.T) {
	tables := []struct {
		val     string
		headers []string
		result  bool
	}{
		{"", []string{}, false},
		{"1, book, 1.5", []string{constants.HeaderQuantity, constants.HeaderProduct, constants.HeaderPrice}, true},
		{"1, book, 1.5", []string{constants.HeaderQuantity, constants.HeaderPrice, constants.HeaderProduct}, false},
		{"1, 1.5, book", []string{constants.HeaderQuantity, constants.HeaderProduct, constants.HeaderPrice}, false},
		{"1.2, book, 1.5", []string{constants.HeaderQuantity, constants.HeaderProduct, constants.HeaderPrice}, false},
		{"1.5, 2, book", []string{constants.HeaderPrice, constants.HeaderQuantity, constants.HeaderProduct}, true},
	}

	for _, table := range tables {
		var it Items
		item, err := it.Save(table.val, table.headers)
		if err != nil && len(item) > 0 && table.result == true {
			t.Errorf("Save(%v, %v) was incorrect, got: %v, want: %v.", table.val, table.headers, err, nil)
			break
		}
		if err == nil && len(item) == 0 {
			t.Errorf("Save(%v, %v) was incorrect, got: %v, want: %v.", table.val, table.headers, item, table.result)
		}
	}
}
