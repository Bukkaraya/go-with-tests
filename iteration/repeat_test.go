package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("n", 10)
	expected := "nnnnnnnnnn"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func ExampleRepeat() {
	repeated := Repeat("x", 2)
	fmt.Println(repeated)
	// Output: xx
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func TestStandardLibRepeat(t *testing.T) {
	repeated := strings.Repeat("a", 3)
	expected := "aaa"

	if expected != repeated {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}
