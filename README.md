# Go Map Operations and Error Handling

This repository contains a Go program that demonstrates common map operations and how to handle potential errors, such as accessing non-existent keys.

## Program Overview

The `bug.go` file contains a Go program that illustrates the following map operations:

- Creating a map
- Inserting key-value pairs
- Accessing elements
- Deleting elements
- Checking for element existence
- Iterating over the map

The program showcases how Go's maps handle missing keys by returning the zero value for the corresponding data type. It also demonstrates the use of the comma ok idiom for checking if a key exists before accessing its value. 

## How to Run

1. Clone this repository.
2. Navigate to the repository's directory in your terminal.
3. Run the program using the command `go run bug.go`.

## Additional Notes

The program's comments explain each step in detail. The program demonstrates best practices for working with maps in Go, ensuring safe and predictable behavior when handling potentially missing keys.