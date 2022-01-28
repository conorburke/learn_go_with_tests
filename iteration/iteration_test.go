package iteration

import (
	"fmt"
	"testing"
)

func TestIteration(t *testing.T) {
    t.Run("should print out n times", func(t *testing.T) {
        expected := "xxxxx"
        got := Iterate("x", 5)

        if expected != got {
            t.Errorf("expected %q but got %q", expected, got)
        }

    })
    
}

func BenchmarkRepeat(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Iterate("x", 5)
    }
}

func ExampleIterate() {
    output := Iterate("z", 10)
    fmt.Println(output)
    // Output: zzzzzzzzzz
}