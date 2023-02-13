package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	logger := log.Default()
	app := NewApplication(logger)
	app.run()
}

type Application struct {
	logger *log.Logger
}

func NewApplication(logger *log.Logger) *Application {
	return &Application{logger: logger}
}

func (a Application) run() {
	// user wants to input via console
	inputStringPtr := flag.String("c", "", "The intervals you want to merge. Written in the following manner: lowerBound,upperBound;lowerBound,upperBound")

	// user wants to input via file
	inputFilePtr := flag.String("f", "", "A file containing new line seperated intervals you want to use as input for the merge")

	flag.Parse()
	// check if input string is existent
	if *inputStringPtr != "" {
		inputString := *inputStringPtr
		// split up input string at all occurring ;
		intervalStrs := strings.Split(inputString, ";")

		// from here on the handling is the same as read in from a file
		a.mergeIntervalStrs(intervalStrs)

	} else if *inputFilePtr != "" {
		// try to open the given file
		file, err := os.Open(*inputFilePtr)

		if err != nil {
			a.logger.Fatalf("Could not read file at given location!")
		}

		// make sure file is closed at the end
		defer file.Close()

		sc := bufio.NewScanner(file)

		var intervalStrs []string

		// scan through the lines step by step
		for sc.Scan() {
			intervalStrs = append(intervalStrs, sc.Text())
		}

		// from here on the handling is the same as with a console input
		a.mergeIntervalStrs(intervalStrs)
	} else {
		a.logger.Fatalf("Provide at least one inputflag!")

	}

}

// takes strings representing intervals by commaseperated upper and lower bound, converts them safely to [Interval] and merges them
func (a Application) mergeIntervalStrs(intervalStrs []string) {
	var intervals []*Interval
	for i, str := range intervalStrs {
		// upper and lower bound are expected to be seperated by ,
		splitStr := strings.Split(str, ",")

		// handle possible bad/malicious inputs
		if len(splitStr) < 2 {
			a.logger.Fatalf("Input String inproperly formatted! You have provided no upper bound for interval at position %d", i)
		}
		if len(splitStr) > 2 {
			a.logger.Fatalf("Input String inproperly formatted! You have provided too many values for interval at position %d", i)
		}

		// convert lower and upper bound to ints and handle possible errors
		lowerBound, err := strconv.Atoi(splitStr[0])
		if err != nil {
			a.logger.Fatalf("Input String inproperly formatted! You have provided a non integer value  for the lower bound of the interval at position %d", i)
		}
		upperBound, err := strconv.Atoi(splitStr[1])
		if err != nil {
			a.logger.Fatalf("Input String inproperly formatted! You have provided a non integer value  for the upper bound of the interval at position %d", i)
		}

		intervals = append(intervals, NewInterval(lowerBound, upperBound))

	}
	// the actual merging
	mergedIntervals := Merge(intervals)

	// print all merged intervals
	for _, interval := range mergedIntervals {
		println(interval.String())
	}
	// we're done here
	os.Exit(0)
}
