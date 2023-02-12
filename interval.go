package main

// Interval is used to represent an interval of integers
type Interval struct {
	LowerBound int
	UpperBound int
}

// NewInterval basic for an interval
func NewInterval(lowerBound int, upperBound int) *Interval {
	// handle possible invalid input
	if !(lowerBound <= upperBound) {
		panic("lowerBound must be smaller or equal to upperbound!")
	}
	return &Interval{LowerBound: lowerBound, UpperBound: upperBound}
}

// we implement the [sort.Interface] here for Interval[]

type ByLowerBound []*Interval

// Len return the length of the slice to be sorted
func (is ByLowerBound) Len() int {
	return len(is)
}

// Swap swaps the elements at given indexes
func (is ByLowerBound) Swap(i, j int) { is[i], is[j] = is[j], is[i] }

// Less evaluates if value at i should be considered less than value at j
func (is ByLowerBound) Less(i, j int) bool { return is[i].LowerBound < is[j].LowerBound }
