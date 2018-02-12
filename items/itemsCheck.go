package items

import "strings"

// isBook checks if product is a book
func (it *Items) isBook(v string) bool {
	keywords := []string{"book"}
	return it.isFound(v, keywords)
}

// isFood checks if a product is a food
func (it *Items) isFood(v string) bool {
	keywords := []string{"choco", "wine", "cake", "wheat", "bread", "cand", "wine", "cookie"}
	return it.isFound(v, keywords)
}

// isFound checks if a keyword is within a collection of specified product
func (it *Items) isFound(v string, keywords []string) bool {
	// Check if any of the keywords are present
	for _, val := range keywords {
		if strings.Contains(strings.ToLower(v), strings.ToLower(val)) {
			return true
		}
	}
	return false
}

// isImported checks if product is imported
func (it *Items) isImported(v string) bool {
	keywords := []string{"import"}
	return it.isFound(v, keywords)
}

// isMedicine checks if a product is a medicine
func (it *Items) isMedicine(v string) bool {
	keywords := []string{"pill", "med"}
	return it.isFound(v, keywords)
}
