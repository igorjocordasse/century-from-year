package main

import (
	"fmt"
)

func centuryFromYear(year int64) int64 {

	result := year / 100
	y := year % 100

		if y == 0 {
			return result
		} else {
			return result + 1
		}
}

func main() {

	r := centuryFromYear(2001)
	fmt.Println("the Century is ", r)
}
