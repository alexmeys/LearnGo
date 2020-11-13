package p193

import (
	"strings"
	"testing"
)

const s = "We ask ourselves, Who am I to be brilliant, gorgeous talented, faboulous? Actually, who are you not to be? Your playing small does not serve the world. There nothing enlightened about shrinking so that other people won't feel insecure around you. We are all meant to shine, as us; it's in everyone. And as we let our own light shine, we unconsciously give other people permission to do the same. As we are liberated from our own fear, our presence automatically liberates others. - Marianne Williamson"

func BenchmarkCat(b *testing.B) {
	xs := strings.Split(s, " ")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Cat(xs)
	}
}
func TestCat(t *testing.T) {
	if s != s {
		t.Error("Got", s, "Expected", s)
	}
}
