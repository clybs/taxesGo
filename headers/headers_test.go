package headers

import (
	"reflect"
	"testing"
)

func TestIsStringInSlice(t *testing.T) {
	var hd Headers
	tables := []struct {
		val    string
		aVal   []string
		result bool
	}{
		{"a", []string{}, false},
		{"a", []string{"a"}, true},
		{"a", []string{"a", "b"}, true},
		{"a", []string{"b", "a"}, true},
		{"a", []string{"b", "c"}, false},
	}

	for _, table := range tables {
		result := hd.isStringInSlice(table.val, table.aVal)
		if result != table.result {
			t.Errorf("isStringInSlice(%v, %v) was incorrect, got: %v, want: %v.", table.val, table.aVal, result, table.result)
		}
	}
}

func TestIsValidHeader(t *testing.T) {
	tables := []struct {
		val    string
		result bool
	}{
		{"a", false},
		{"a, b, c", false},
		{"product, price, c", false},
		{"product, price, quantity", true},
		{"proDuct , Price, quanTity   ", true},
		{"price, product,   quantity", true},
		{",,,,", false},
		{"product price quantity", false},
	}

	for _, table := range tables {
		var hd Headers
		valid, entries := hd.IsValid(table.val)
		if valid != table.result {
			t.Errorf("IsValid(%v) was incorrect, got: %v, want: %v.", table.val, valid, table.result)
			break
		}
		if valid && len(entries) == 0 {
			t.Errorf("IsValid(%v) was incorrect, got: %v, want: %v.", table.val, len(entries), "count > 0")
		}
	}
}

func TestSaveResetOrder(t *testing.T) {
	var hd Headers
	tables := []struct {
		Save    bool
		headers []string
		result  []string
	}{
		{true, []string{"a"}, []string{"a"}},
		{true, []string{"a", "b"}, []string{"a", "b"}},
		{true, []string{"b", "c"}, []string{"b", "c"}},
		{false, []string{"a"}, []string{}},
		{false, []string{"a", "b"}, []string{}},
	}

	for _, table := range tables {
		hd.SaveResetOrder(table.Save, table.headers)
		result := hd.Values
		if !reflect.DeepEqual(result, table.result) {
			t.Errorf("SaveResetOrder(%v, %v) was incorrect, got: %v, want: %v.", table.Save, table.headers, result, table.result)
		}
	}
}
