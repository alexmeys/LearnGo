package p191

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	s := Greet("Jonas")
	if s != "Welcome Jonas" {
		t.Error("Got", s, "Expected", "Welcome Jonas")
	}
}

func ExampleGreet() {
	fmt.Println(Greet("Jonas"))
	// Output:
	// Welcome Jonas
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("Jonas")
	}
}
