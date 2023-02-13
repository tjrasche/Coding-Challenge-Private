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
./merge "input string"
```
You may also directly run the program with 

```bash
go run main.go "input string"
```

## Inputting range values

The program expects exactly one argument, a string representing all intervals that shall be merged.

A single interval is written as teh lower and upper bound seperated by a comma.
Multiple intervals may be seperated by a semicolon.

So to input the example values given in the coding task would be formatted as follows:

```
"25,30;2,19;14,23;4,8"
```

# Considerations regarding the Runtime and Memory 
The algorithms runtime has an average time complexity of O(n*logn). 
This stems from the following considerations:
    Complexity of sorting: O(n*logn)
    Complexity of merging for loop: O(n)

The algorithms auxiliary space complexity is currently O(n),
as the algorithm is writing the merged values into a new slice.

This could be omitted by making the algorithm in place.

# Known limitations
The program can currently only handle integers.
The program is limited to input sizes that can be handled in memory.