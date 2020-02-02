package custompkg

import (
	"fmt"
	"testing"
)

func TestCustomsum(t *testing.T) {
	type testcase struct {
		data []int
		out  int
	}
	// data for table tests
	testcases := []testcase{
		testcase{[]int{1, -1}, 0},
		testcase{[]int{1, 2}, 3},
		testcase{[]int{1, 2, 3}, 6},
		testcase{[]int{1, 2, 3, 4}, 10},
		testcase{[]int{1, 2, 3, 4, 5}, 15},
	}

	for _, tc := range testcases {
		out := customsum(tc.data)
		if out != tc.out {
			t.Error("Expected", tc.out, "Got:", out)
		}
	}
}

func TestCustomAvg(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	v := CustomAvg(s)
	if v != 3 {
		t.Error("Expected 3, got", v)
	}
}

func ExampleCustomAvg() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(CustomAvg(s))
	// Output:
	// 3
}
