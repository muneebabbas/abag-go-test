package main

func main() {
}

// isAestheticallyPleasing checks if the sequence is aesthetically pleasing
// aesthetically pleasing is defined if the numbers alternately increase
// and decrease. The start could be increasing or decreasing
func isAestheticallyPleasing(arr []int) bool {
	isSequenceDecreasing := true
	if arr[0] > arr[1] {
		isSequenceDecreasing = false
	}

	for i := 0; i < len(arr)-1; i++ {
		if isSequenceDecreasing {
			if !(arr[i] < arr[i+1]) {
				return false
			}
		} else {
			if !(arr[i] > arr[i+1]) {
				return false
			}
		}
		isSequenceDecreasing = !isSequenceDecreasing
	}
	return true
}

// Solution returns the number of ways of cutting one element from
// the array so that the resulting array remains aesthetically pleasing
// it just cuts each element and checks the resulting array
func Solution(A []int) int {
	if isAestheticallyPleasing(A) {
		return 0
	}
	counter := 0
	for i := range A {
		newA := make([]int, len(A))
		copy(newA, A)
		cutArr := remove(newA, i)
		if isAestheticallyPleasing(cutArr) {
			counter++
		}
	}
	if counter == 0 {
		return -1
	}
	return counter
}

// remove element from slice
func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

// OptimizedSolution tries to make sure that at every point there is either a peak
// or a valley and they alternate between each other.
// peak = a < b > c
// valley = a > b < c
// If the desired condition does not hold, we try to cut a, b and c
// and check which ones result in a peak -> valley or valley -> peak
func OptimizedSolution(A []int) int {
	isSequenceDecreasing := true

	if A[0] == A[1] {
		if A[0] < A[2] {
			isSequenceDecreasing = false
		}
	} else if A[0] < A[1] {
		isSequenceDecreasing = false
	}

	cutsMade := false
	count := 0
	for i := 0; i < len(A)-2; i++ {
		a := A[i]
		b := A[i+1]
		c := A[i+2]
		// If Sequence is decreasing, we should have a valley
		// Check if we have a valid valley
		if isSequenceDecreasing {
			if !(a > b && b < c) {
				if cutsMade {
					return -1
				}
				// cut a
				// a-1, b, c should be a peak
				// b, c, c+1 should be valley
				if isPeak(i-1, i+1, i+2, A) && isValley(i+1, i+2, i+3, A) {
					count++
				}

				// cut b
				// a-1, a, c should be a peak
				// a, c, c+1 should be valley
				if isPeak(i-1, i, i+2, A) && isValley(i, i+2, i+3, A) {
					count++
				}

				// cut c
				// a-1, a, b should be a peak
				// a, b, c+1 should be valley
				if isPeak(i-1, i, i+1, A) && isValley(i, i+1, i+3, A) {
					count++
				}

				cutsMade = true
				continue
			}

		}
		if !isSequenceDecreasing {
			if !(a < b && b > c) {
				if cutsMade {
					return -1
				}
				// cut a
				// a-1, b, c should be a valley
				// b, c, c+1 should be peak
				if isValley(i-1, i+1, i+2, A) && isPeak(i+1, i+2, i+3, A) {
					count++
				}

				// cut b
				// a-1, a, c should be a valley
				// a, c, c+1 should be peak
				if isValley(i-1, i, i+2, A) && isPeak(i, i+2, i+3, A) {
					count++
				}

				// cut c
				// a-1, a, b should be a valley
				// a, b, c+1 should be peak
				if isValley(i-1, i, i+1, A) && isPeak(i, i+1, i+3, A) {
					count++
				}
				cutsMade = true
				continue
			}
		}
		isSequenceDecreasing = !isSequenceDecreasing

	}
	return count
}

// isValley checks i > j && j < k
func isValley(i, j, k int, arr []int) bool {
	if inRange(i, arr) && inRange(j, arr) {
		if arr[i] <= arr[j] {
			return false
		}
	}
	if inRange(j, arr) && inRange(k, arr) {
		if arr[j] >= arr[k] {
			return false
		}
	}
	return true
}

// isPeak checks i < j and j > k
func isPeak(i, j, k int, arr []int) bool {
	if inRange(i, arr) && inRange(j, arr) {
		if arr[i] >= arr[j] {
			return false
		}
	}
	if inRange(j, arr) && inRange(k, arr) {
		if arr[j] <= arr[k] {
			return false
		}
	}
	return true
}

func inRange(i int, arr []int) bool {
	if i >= 0 && i < len(arr) {
		return true
	}
	return false
}
