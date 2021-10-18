package main

import "testing"

func TestIsAestheticallyPleasing(t *testing.T) {
	arr := []int{1, 10, 4, 100, -10, 50, 5, 100}
	if isAestheticallyPleasing(arr) != true {
		t.Errorf("isAestheticallyPleasing with %v should be true", arr)
	}
	arr = []int{5, 10, 5, 10, 5, 10}
	if isAestheticallyPleasing(arr) != true {
		t.Errorf("isAestheticallyPleasing with %v should be true", arr)
	}

	arr = []int{1, 10, 15, 10, 5, 10}
	if isAestheticallyPleasing(arr) != false {
		t.Errorf("isAestheticallyPleasing with %v should be false", arr)
	}
	arr = []int{10, 5, 10}
	if isAestheticallyPleasing(arr) != true {
		t.Errorf("isAestheticallyPleasing with %v should be true", arr)
	}
	arr = []int{10, 5, 10, 5, 4, 3}
	if isAestheticallyPleasing(arr) != false {
		t.Errorf("isAestheticallyPleasing with %v should be false", arr)
	}
}

func TestIsValley(t *testing.T) {
	if isValley(0, 1, 2, []int{5, 4, 5}) != true {
		t.Errorf("Wanted true, got false")
	}
	if isValley(0, 1, 2, []int{5, 4, 4}) != false {
		t.Errorf("Wanted false, got true")
	}
	if isValley(0, 1, 2, []int{4, 5, 4}) != false {
		t.Errorf("Wanted false, got true")
	}
	if isValley(0, 1, 2, []int{4, 4, 4}) != false {
		t.Errorf("Wanted false, got true")
	}
}

func TestIsPeak(t *testing.T) {
	if isPeak(0, 1, 2, []int{4, 5, 4}) != true {
		t.Errorf("Wanted true, got false")
	}
	if isPeak(0, 1, 2, []int{5, 4, 4}) != false {
		t.Errorf("Wanted false, got true")
	}
	if isPeak(0, 1, 2, []int{4, 3, 4}) != false {
		t.Errorf("Wanted false, got true")
	}
	if isPeak(-1, 1, 2, []int{4, 4, 4}) != false {
		t.Errorf("Wanted false, got true")
	}
	if isPeak(0, 1, 2, []int{1, 4, 5}) != false {
		t.Errorf("Wanted false, got true")
	}
}

func TestSolution(t *testing.T) {

	type errorTestCase struct {
		input  []int
		result int
	}

	testCases := []errorTestCase{
		{
			input:  []int{3, 4, 5, 3, 7},
			result: 3,
		},
		{
			input:  []int{1, 3, 1, 2},
			result: 0,
		},
		{
			input:  []int{1, 2, 3, 4, 5, 6, 7},
			result: -1,
		},
		{
			input:  []int{10, 5, 11, 12, 3, 5, 2},
			result: 2,
		},
		{
			input:  []int{10, 5, 6, 4, 5, 2, 2},
			result: 2,
		},
		{
			input:  []int{5, 5, 5, 4, 5, 2, 2},
			result: -1,
		},
		{
			input:  []int{1, 700, 200, 356, 487, 333, 350, 200},
			result: 3,
		},
		{
			input:  []int{5, 10, 5, 10, 5, 10, 5, 10, 5, 10},
			result: 0,
		},
		{
			input:  []int{5, 4, 3, 6, 2, 10},
			result: 3,
		},
		{
			input:  []int{5, 5, 2, 10, 1, 100},
			result: 2,
		},
		{
			input:  []int{5, 5, 10, 1, 100, -20},
			result: 2,
		},
	}

	for _, testCase := range testCases {
		optimizedSol := OptimizedSolution(testCase.input)
		sol := Solution(testCase.input)

		if optimizedSol != testCase.result {
			t.Errorf("Array: %v - Expected %d, Solution (Optimized): %d", testCase.input, testCase.result, optimizedSol)
		}
		if sol != testCase.result {
			t.Errorf("Array: %v - Expected %d, Solution (BruteForce): %d", testCase.input, testCase.result, sol)
		}
	}

}
