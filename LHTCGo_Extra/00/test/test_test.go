package test

import "testing"

func BenchmarkUp(b *testing.B) {
	for i:= 0; i<b.N; i++{
		Up("user")
	}
}

func BenchmarkCaps(b *testing.B){
	for i:= 0; i<b.N; i++{
		Caps("alex")
	}
}