// Package sample is used to show a sample
package sample

import (
	"fmt"
	"testing"
)

func TestSayHi(t *testing.T) {
	s := SayHi("Mohit")
	if s != "Hi Mohit" {
		t.Error("Expected: Hi Mohit", "got:",s)
	}
}

func ExampleSayHi() {
	fmt.Println(SayHi("Mohit"))
	// Output:
	// Hi Mohit
}

func BenchmarkSayHi(b *testing.B){
	for i := 0 ; i < b.N ; i ++ {
		SayHi("Mohit")
	}
}
