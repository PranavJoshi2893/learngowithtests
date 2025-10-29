package integers

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	got := Sum(10, 12)
	want := 22

	if got != want {
		t.Errorf("got: %d want: %d", got, want)
	}
}

func ExampleSum() {
	sum := Sum(2, 4)
	fmt.Println(sum)
	//output: 6
}
