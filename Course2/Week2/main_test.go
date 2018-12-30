package main

import "testing"

func TestComputeDisplacement(t *testing.T) {
	testcases := []struct {
		time         float64
		displacement float64
	}{
		{3, 52},
		{5, 136},
		{0, 1},
	}

	fn := GenDisplaceFn(10, 2, 1)

	for _, testcase := range testcases {
		result := fn(testcase.time)
		if result != testcase.displacement {
			t.Errorf("Displacement computation went wrong for 10, 2, 1 and %f time, got: %f, want: %f", testcase.time, result, testcase.displacement)
		}
	}
}
