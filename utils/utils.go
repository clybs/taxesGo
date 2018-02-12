package utils

import "strings"

type Utils struct{}

// GetEntries gets the entries
func (ut *Utils) GetEntries(v string) []string {
	entries := strings.Split(strings.ToLower(v), ",")

	// Remove the white space
	entriesTrimmed := make([]string, 0)
	for _, v := range entries {
		val := strings.Trim(v, " ")

		// Get only entries with values
		if val != "" {
			entriesTrimmed = append(entriesTrimmed, val)
		}
	}

	return entriesTrimmed
}
