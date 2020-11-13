package p196

import (
	"fmt"
	"testing"
)

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count("The string")
	}
}

func TestCount(t *testing.T) {
	a := Count("The string")
	if a != 2 {
		t.Error("Got", a, "Want", 2)
	}
}

func ExampleCount() {
	fmt.Println(Count("The String"))
	// Output:
	// 2
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count("The string has multiple string values.")
	}
}
