package main

import (
	"math/rand"
	"testing"
	"time"
)

// This is a basic first test using the values provided in teh coding challenge to test the merge algorithm
func TestMergeExample(t *testing.T) {

	intervals := []*Interval{NewInterval(25, 30), NewInterval(2, 19), NewInterval(14, 23), NewInterval(4, 8)}

	mergedIntervals := Merge(intervals)

	mergedIntervalsExpectedResult := []*Interval{NewInterval(2, 23), NewInterval(25, 30)}

	// check if results from Merge() are equivalent to expected case
	if len(mergedIntervals) != len(mergedIntervalsExpectedResult) {
		t.Errorf("Expected lengths of result to be equal")
	}

	for i, interval := range mergedIntervalsExpectedResult {
		t.Log(*interval)
		t.Log(*mergedIntervals[i])
		if !(*interval == *mergedIntervals[i]) {
			t.Errorf("Expected Values to be equivalent and in order! Expected value: %s, value reveived: %s", *interval, *mergedIntervals[i])
		}
	}
}

// Test to measure performance of the sorting. Consecetuvely merges larger and larger values and measures the time they take.
// this is done, so a tester can roughly evaluate the growth rate of the execution time of the merge algorithm.
func TestPerformance(t *testing.T) {
	// slowly let the result size grow
	interValCounts := []int{10, 100, 1000, 10_000, 100_000, 1_000_000}
	var results []time.Duration
	for _, interValCount := range interValCounts {
		results = append(results, timedMergeWithRandomIntervals(t, interValCount))
	}
	var execTimeQuotients []float64
	for i, duration := range results {
		execTimeQuotients = append(execTimeQuotients, float64(duration.Nanoseconds())/float64(interValCounts[i]))

		// we expect a runtime behaviour of O(n*log(n)), so we can fail if the time per interval does not drop with bigger intervals
		if i != 0 && (execTimeQuotients[i-1] < execTimeQuotients[i]) {
			t.Errorf("Runtime scaling worse than linear")
		}
	}
}

// generates n random intervals and merges them. Returns the execution time of the merge itself.
func timedMergeWithRandomIntervals(t *testing.T, n int) (execTime time.Duration) {
	// generate a very large slice of intervals
	rand.Seed(time.Now().UnixNano())

	const intervalsCount = 10_000_000
	var intervals []*Interval

	for i := 0; i < intervalsCount; i++ {
		lowerBound := rand.Intn(100000)
		upperBound := lowerBound + rand.Intn(100000)
		intervals = append(intervals, NewInterval(lowerBound, upperBound))
	}

	// only measure the time the actual merging takes. Not the random generation of intervals
	startTime := time.Now()
	Merge(intervals)
	return time.Now().Sub(startTime)
}
