// Package p189 used for test app
package p189

// Sum functin takes unlimited ints, and returns an int
func Sum(xi ...int) int {
	s := 0
	for _, v := range xi {
		s += v
	}
	return s
}
