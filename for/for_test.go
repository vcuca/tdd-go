package iteration

import (
	"fmt"
	"testing"
)

func ExampleRepeat() {
	repeated := Repeat("a", 2)
	fmt.Println(repeated)
	// Output: aa
}

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 3)
	expected := "aaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 2)
	}
}
