// Package p193 for benchmark tests
package p193

import "strings"

// Cat is a function...
func Cat(xs []string) string {
	s := ""
	for _, v := range xs {
		s += v
		s += " "
	}
	return s
}

// Join is a function...
func Join(xs []string) string {
	return strings.Join(xs, " ")
}
