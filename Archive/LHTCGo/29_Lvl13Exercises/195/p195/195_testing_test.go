package p195

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	a := Years(10)
	if a != 70 {
		t.Error("Got:", a, "Want:", 70)
	}
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(10)
	}
}

func ExampleYears() {
	fmt.Println(Years(10))
	// Output:
	// 70
}
