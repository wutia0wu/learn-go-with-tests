package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("T", 5)
	expected := "TTTTT"

	if repeated != expected {
		t.Errorf("expected '%q' but got '%q'", expected, repeated)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 3)
	fmt.Println(repeated)
	// Output: aaa
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("T", b.N)
	}
}
