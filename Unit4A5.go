// Author: William Provost
// Version: 1.0.0
// Date: 2025-11-25
// Fileoverview: This program will print out a times table for the values from 1 to 12, for each of the 12 values in the table.

package main

import (
	"fmt"
)

func main() {
	// variables
	var row string

	// output
	fmt.Println("Times Table from 1 to 12:")

	// Loop through the numbers 1 to 12 for the rows
	for rowVariable := 1; rowVariable <= 12; rowVariable++ {
		// Loop through the numbers 1 to 12 for the columns
		for columnVariable := 1; columnVariable <= 12; columnVariable++ {
			// Calculate the product and format it to align properly
			// fmt.Sprintf("%4d", ...) formats the integer to take up 4 characters
			row += fmt.Sprintf("%4d", rowVariable*columnVariable)
		}
		// Print the completed row
		fmt.Println(row)
		// Reset the row variable for the next iteration
		row = ""
	}

	fmt.Println("\nDone.")
}
