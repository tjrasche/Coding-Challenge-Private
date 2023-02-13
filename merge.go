package main

import "sort"

// Merge merges all [Interval] together when they have overlapping values, returns values sorted by LowerBound of the Intervals
func Merge(intervals ByLowerBound) []*Interval {
	// first we have to sort all intervals. this is done using the [sort.Sort] method provided by go.
	// sort.Sort uses the methods defined for [ByLowerBound] and described by the [sort.Intercace] to sort the values in an ascending matter
	sort.Sort(intervals)

	// used for storing the value of the latest "mergetarget" for all intervals
	mergeTargetIndex := 0

	for _, unmergedInterval := range intervals {

		// change < to <= if you want to consider two intervals x,y to be overlapping if x.UpperBound = y.LowerBound
		if unmergedInterval.LowerBound < intervals[mergeTargetIndex].UpperBound {
			// the following interval is overlapping with the current mergetarget
			// we have to enlarge the mergeTarget interval,
			// if the upperBound of the unmergedInterval is larger than the upperbound of the current mergeTarget
			if unmergedInterval.UpperBound > intervals[mergeTargetIndex].UpperBound {
				intervals[mergeTargetIndex].UpperBound = unmergedInterval.UpperBound
			}
		} else {
			// we don't have any overlap more. so the currentInterval can be considered "done" and we can step up the mergeTargetIndex
			// and set the value at the new mergeTargetIndex to the current interval as starting point
			mergeTargetIndex++
			intervals[mergeTargetIndex] = unmergedInterval
		}
	}
	return intervals[:mergeTargetIndex+1]
}
