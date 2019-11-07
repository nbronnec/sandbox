package main

import (
	"fmt"
	"os"
)

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	var N int
	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		var operation, arg1, arg2 string
		fmt.Scan(&operation, &arg1, &arg2)

		fmt.Fprintln(os.Stderr, operation, " ", arg1, " ", arg2)

	}
	for i := 0; i < N; i++ {

		fmt.Fprintln(os.Stderr, "Debug messages...")
		fmt.Println("1") // Write answer to stdout
	}
}
