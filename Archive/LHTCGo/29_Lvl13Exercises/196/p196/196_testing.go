// Package p196 gives wordcount answers
package p196

import "strings"

// UseCount counts the words that are similar
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Count counts the words
func Count(s string) int {
	xs := strings.Fields(s)
	return len(xs)
}
