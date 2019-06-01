// Problem link:
// https://www.hackerearth.com/practice/basic-programming/input-output/basics-of-input-output/practice-problems/algorithm/find-factorial/

package main

import (
	"fmt"
)

func main() {
	var l int
	fmt.Scan(&l)
	answer := 1
	for i := 2; i <= l; i++ {
		answer *= i
	}
	fmt.Println(answer)
}
