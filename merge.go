package main

import "sort"

// Merge merges all [Interval] together, wehen they have overlapping values, returns values sorted by LowerBound of the Intervals
func Merge(intervals ByLowerBound) []*Interval {
	// first we have to sort all intervals. this is done using the [sort.Sort] method provided by go.
	// sort.Sort uses the methods defined for [ByLowerBound] and described by the [sort.Intercace] to sort the values in an ascending matter
	sort.Sort(intervals)
	// used for storing the current interval being built
	var currentInterval = Interval{}
	var mergedIntervals []*Interval

	for i, interval := range intervals {

		if i == 0 {
			// we are at the start of the for loop so we have to set the first upper and lower bound and can move on
			currentInterval = *NewInterval(interval.LowerBound, interval.UpperBound)
			continue
		}
		if interval.LowerBound < currentInterval.UpperBound {
			// the following interval is overlapping with the previous one, so we have to enlargen our currentInterval, if the upperBound of the interval is larger than the currentinterval
			if interval.UpperBound > currentInterval.UpperBound {
				currentInterval.UpperBound = interval.UpperBound
			}
		} else {
			// we don't have any overlap more. so the currentInterval can be considered "done" and appendend to the resultSLice
			mergedIntervals = append(mergedIntervals, NewInterval(currentInterval.LowerBound, currentInterval.UpperBound))
			// as we are a starting a new currentInterval, we have to set it's
			currentInterval = *NewInterval(interval.LowerBound, interval.UpperBound)
		}
	}
	// we still have to add the last interval
	mergedIntervals = append(mergedIntervals, NewInterval(currentInterval.LowerBound, currentInterval.UpperBound))
	return mergedIntervals
}
