package utils

import (
	"reflect"
	"testing"
)

func TestGetEntries(t *testing.T) {
	var ut Utils
	tables := []struct {
		val    string
		result []string
	}{
		{"", []string{}},
		{"a", []string{"a"}},
		{"a, b", []string{"a", "b"}},
		{"a , b   , c,", []string{"a", "b", "c"}},
	}

	for _, table := range tables {
		result := ut.GetEntries(table.val)
		if !reflect.DeepEqual(result, table.result) {
			t.Errorf("getEntries(%v) was incorrect, got: %v, want: %v.", table.val, result, table.result)
		}
	}
}
