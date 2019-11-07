package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func order(x, y int) (low, high int) {
	if x > y {
		low = y
		high = x
	} else {
		low = x
		high = y
	}
	return
}

func main() {
	fmt.Println(add(42, 13))
}
