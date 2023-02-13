package main

import (
	"math/rand"
	"testing"
	"time"
)

// This is a basic first test using the values provided in teh coding challenge to test the merge algorithm
func TestMergeExample(t *testing.T) {

	intervals := []*Interval{NewInterval(25, 30), NewInterval(2, 19), NewInterval(14, 23), NewInterval(4, 8)}
	mergedIntervalsExpectedResult := []*Interval{NewInterval(2, 23), NewInterval(25, 30)}

	mergedIntervals := Merge(intervals)

	// check if results from Merge() are equivalent to expected case
	if len(mergedIntervals) != len(mergedIntervalsExpectedResult) {
		t.Errorf("Expected lengths of result to be equal")
	}

	for i, interval := range mergedIntervalsExpectedResult {
		if !(*interval == *mergedIntervals[i]) {
			t.Errorf("Expected Values to be equivalent and in order! Expected value: %s, value reveived: %s", *interval, *mergedIntervals[i])
		}
	}
}

// Test if ranges of negative integers are merged correctly
func TestMergeNegativeNumbers(t *testing.T) {
	intervals := []*Interval{NewInterval(-100, -20), NewInterval(-50, -5), NewInterval(0, 20)}
	mergedIntervalsExpectedResult := []*Interval{NewInterval(-100, -5), NewInterval(0, 20)}
	mergedIntervals := Merge(intervals)

	// check if results from Merge() are equivalent to expected case
	if len(mergedIntervals) != len(mergedIntervalsExpectedResult) {
		t.Errorf("Expected lengths of result to be equal")
	}

	for i, interval := range mergedIntervalsExpectedResult {
		if !(*interval == *mergedIntervals[i]) {
			t.Errorf("Expected Values to be equivalent and in order! Expected value: %s, value reveived: %s", *interval, *mergedIntervals[i])
		}
	}
}

// Test if intervals ranging from negative to positive integers are merged correctly
func TestMergeNegativeAndPositiveNumbers(t *testing.T) {
	intervals := []*Interval{NewInterval(-100, -20), NewInterval(-50, -5), NewInterval(-6, 1), NewInterval(2, 5)}
	mergedIntervalsExpectedResult := []*Interval{NewInterval(-100, 1), NewInterval(2, 5)}
	mergedIntervals := Merge(intervals)

	// check if results from Merge() are equivalent to expected case
	if len(mergedIntervals) != len(mergedIntervalsExpectedResult) {
		t.Errorf("Expected lengths of result to be equal")
	}

	for i, interval := range mergedIntervalsExpectedResult {
		if !(*interval == *mergedIntervals[i]) {
			t.Errorf("Expected Values to be equivalent and in order! Expected value: %s, value reveived: %s", *interval, *mergedIntervals[i])
		}
	}
}

// Test if intervals with same start and end number are considered equal intervals
func TestMergeWithSameUpperLowerBound(t *testing.T) {
	intervals := []*Interval{NewInterval(-100, -20), NewInterval(-20, -5)}
	mergedIntervalsExpectedResult := []*Interval{NewInterval(-100, -20), NewInterval(-20, -5)}
	mergedIntervals := Merge(intervals)

	// check if results from Merge() are equivalent to expected case
	if len(mergedIntervals) != len(mergedIntervalsExpectedResult) {
		t.Errorf("Expected lengths of result to be equal")
	}

	for i, interval := range mergedIntervalsExpectedResult {
		if !(*interval == *mergedIntervals[i]) {
			t.Errorf("Expected Values to be equivalent and in order! Expected value: %s, value reveived: %s", *interval, *mergedIntervals[i])
		}
	}
}

// Test to measure performance of the sorting. Consecutively merges larger and larger values and measures the time they take.
// this is done, so a tester can roughly evaluate the growth rate of the execution time of the merge algorithm.
func TestPerformance(t *testing.T) {
	// slowly let the result size grow
	interValCounts := []int{10, 100, 1000, 10_000, 100_000, 1_000_000}
	// how many times shall the merging be done per interValCount?
	numberOfRepetitions := 50
	var results []float64
	for i, interValCount := range interValCounts {
		var resultSum float64
		// make the same call multiple times to get a realistic estimation of runtime (sorting runtime is dependent on input ordering, which is random)
		for j := 1; j < numberOfRepetitions; j++ {
			resultSum = resultSum + float64(timedMergeWithRandomIntervals(interValCount).Nanoseconds())/float64(interValCount)
		}

		// get average execution time per interval and add it to the list of results
		avgRuntimePerInterval := resultSum / float64(numberOfRepetitions)
		t.Logf("Average Runtime per Interval: %f", avgRuntimePerInterval)

		// expect runtime performance to increase less than quadratic (rough estimation, our algorithm should perform slightly worse than linear (O(n*log n)
		if i != 0 {
			if results[i-1]*2 <= avgRuntimePerInterval {
				t.Errorf("Runtime Performance declines to fast! ")
			}
		}

		results = append(results, avgRuntimePerInterval)

	}

}

// generates n random intervals and merges them. Returns the execution time of the merge itself.
func timedMergeWithRandomIntervals(n int) (execTime time.Duration) {
	// generate a very large slice of intervals

	var intervals []*Interval

	for i := 0; i < n; i++ {
		lowerBound := rand.Intn(100000)
		upperBound := lowerBound + rand.Intn(100000)
		intervals = append(intervals, NewInterval(lowerBound, upperBound))
	}

	// only measure the time the actual merging takes. Not the random generation of intervals
	startTime := time.Now()
	Merge(intervals)
	return time.Since(startTime)
}
