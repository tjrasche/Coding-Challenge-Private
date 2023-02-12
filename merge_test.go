package main

import "testing"

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
		if *interval == *mergedIntervals[i] {
			t.Errorf("Expected Values to be equivalent and in order! Expected value: %s, value reveived: %s", *interval, *mergedIntervals[i])
		}
	}
}
