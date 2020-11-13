package main

import (
	"testing"
)

func main() {

}

func TestAbc(t *testing.T) {
	// TestAbc
	got := Abc()
	if got != 1 {
		t.Errorf("Abc(-1) = %d; want 1", got)
	}
}

func Abc() int {
	return -1
}
