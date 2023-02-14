# Coding Challenge Private
## Build and run
This software is written in Go.

To build it you must have a current version of [golang](https://go.dev/) installed and run

```bash
go build 
```
in the git root directory.

Then run the program with 
```bash
./merge -c "input string"
```
You may also directly run the program with 

```bash
go run main.go -c "input string"
```

## Inputting range values

The program accepts two type of inputs:

1. direct command line input using the `-c` flag
2. a filepath to a file containing a list of intervals with the `-f` flag

### Inputting vie command line
When inputting via command line, a  single interval is written as the lower and upper bound seperated by a comma.
Multiple intervals may be seperated by a semicolon.

So to input the example values given in the coding task the string would be formatted as follows:

```
"25,30;2,19;14,23;4,8"
```

This string must be used as the value for the -c flag.

Example:

```bash
./merge  -c "25,30;2,19;14,23;4,8"
```
or 
```bash
go run . -c "25,30;2,19;14,23;4,8"
```

### Inputting via file
To use a file as input for the merging simply use the -f flag followed by a file path.

The file must contain one interval per line.

A single interval is written as the lower and upper bound seperated by a comma.

An example file can be found at `examplevals`.

Example:
```bash
./merge  -f examplevals
```
or 
```bash
go run . -f examplevals
```
## Considerations regarding the Runtime and Memory complexity
The algorithms runtime has an average time complexity of O(n*logn). 
This stems from the following considerations:
    Complexity of sorting: O(n*logn)
    Complexity of merging for loop: O(n)

The algorithms auxiliary space complexity is O(1).

## Known limitations
The program can currently only handle integers.
The program is limited to input sizes that can be handled in memory.

## Assumptions taken
Intervals that share the same upper and lower bound are *not* considered as overlapping.

i.e. 25-30 and 30-35 will not be merged.

This can simply be changed by changing the implementation of Merge() in merge.go in line 17 as documented.