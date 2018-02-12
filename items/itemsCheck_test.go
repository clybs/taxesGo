package items

import (
	"testing"
)

func TestIsBook(t *testing.T) {
	var it Items
	tables := []struct {
		val    string
		result bool
	}{
		{"", false},
		{"A New Book", true},
	}

	for _, table := range tables {
		result := it.isBook(table.val)
		if result != table.result {
			t.Errorf("isBook(%v) was incorrect, got: %v, want: %v.", table.val, result, table.result)
		}
	}
}

func TestIsFood(t *testing.T) {
	var it Items
	tables := []struct {
		val    string
		result bool
	}{
		{"", false},
		{"chocolate slice", true},
		{"birthday cake", true},
		{"candies and cookies", true},
		{"bottle of wine", true},
		{"box of shoes", false},
	}

	for _, table := range tables {
		result := it.isFood(table.val)
		if result != table.result {
			t.Errorf("isFood(%v) was incorrect, got: %v, want: %v.", table.val, result, table.result)
		}
	}
}

func TestIsFound(t *testing.T) {
	var it Items
	tables := []struct {
		val     string
		keyword []string
		result  bool
	}{
		{"", []string{""}, true},
		{"A New Book", []string{"book"}, true},
	}

	for _, table := range tables {
		result := it.isFound(table.val, table.keyword)
		if result != table.result {
			t.Errorf("isFound(%v, %v) was incorrect, got: %v, want: %v.", table.val, table.keyword, result, table.result)
		}
	}
}

func TestIsImported(t *testing.T) {
	var it Items
	tables := []struct {
		val    string
		result bool
	}{
		{"", false},
		{"chocolate slice", false},
		{"imported birthday cake", true},
		{"candies and cookies", false},
		{"imported bottle of wine", true},
		{"box of shoes", false},
	}

	for _, table := range tables {
		result := it.isImported(table.val)
		if result != table.result {
			t.Errorf("isImported(%v) was incorrect, got: %v, want: %v.", table.val, result, table.result)
		}
	}
}

func TestIsMedicine(t *testing.T) {
	var it Items
	tables := []struct {
		val    string
		result bool
	}{
		{"", false},
		{"box of pills", true},
		{"medicinal herbs", true},
		{"candies and cookies", false},
		{"bottle of wine", false},
		{"box of shoes", false},
	}

	for _, table := range tables {
		result := it.isMedicine(table.val)
		if result != table.result {
			t.Errorf("isFood(%v) was incorrect, got: %v, want: %v.", table.val, result, table.result)
		}
	}
}
