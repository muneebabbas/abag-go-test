
## Description

This repo provides solution for the provided question.

The main.go files contains a brute force solution which just removes each element from the array and checks if the resulting array is Aesthetically Pleasing.

There is an optimized solution as well which solves the problem in a single pass by finding points where the desired condition does not hold and making relevant cuts.

## Testing

the main_test.go files contains tests that try to cover all of the paths. Run `go test` to run the tests.
