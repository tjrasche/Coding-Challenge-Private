package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	logger := log.Default()

	// check if input string is existent
	if len(os.Args) == 2 {
		inputString := os.Args[1]
		// split up input string at all occurring ;
		intervalStrs := strings.Split(inputString, ";")

		// interval inputs are expected to be seperated by ;
		var intervals []*Interval
		for i, str := range intervalStrs {
			// upper and lower bound are expected to be seperated by ,
			splitStr := strings.Split(str, ",")

			// handle possible bad/malicious inputs
			if len(splitStr) < 2 {
				logger.Fatalf("Input String inproperly formatted! You have provided no upper bound for interval at position %d", i)
			}
			if len(splitStr) > 2 {
				logger.Fatalf("Input String inproperly formatted! You have provided too many values for interval at position %d", i)
			}

			// convert lower and upper bound to ints and handle possible errors
			lowerBound, err := strconv.Atoi(splitStr[0])
			if err != nil {
				logger.Fatalf("Input String inproperly formatted! You have provided a non integer value  for the lower bound of the interval at position %d", i)
			}
			upperBound, err := strconv.Atoi(splitStr[1])
			if err != nil {
				logger.Fatalf("Input String inproperly formatted! You have provided a non integer value  for the upper bound of the interval at position %d", i)
			}

			intervals = append(intervals, NewInterval(lowerBound, upperBound))

		}
		mergedIntervals := Merge(intervals)
		for _, interval := range mergedIntervals {
			println(interval.String())
		}

	} else {
		logger.Fatalf("Invalid number of arguments")
	}
}
