/**
 * @author William Provost
 * @version 1.0.0
 * @date 2025-11-25
 * @fileoverview This program will print out a times table for the values from 1 to 12.
 */

// variables
let row = "";

// output
console.log("Times Table from 1 to 12:");

// Loop through the numbers 1 to 12 for the rows
for (let rowVariable = 1; rowVariable <= 12; rowVariable++) {
  // Loop through the numbers 1 to 12 for the columns
  for (let columnVariable = 1; columnVariable <= 12; columnVariable++) {
    // Calculate the product and format it to align properly
    // .padStart(4, ' ') ensures each number takes up 4 characters for alignment
    row += (rowVariable * columnVariable).toString().padStart(4, ' ');
  }
  // Print the completed row
  console.log(row);
  // Reset the row variable for the next iteration
  row = "";
}

console.log("\nDone.");