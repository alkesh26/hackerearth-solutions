// Problem link:
// https://www.hackerearth.com/practice/basic-programming/input-output/basics-of-input-output/practice-problems/algorithm/find-product/

package main

import (
	"fmt"
)

func main() {
	var size int
	var num int
	answer := 1
	fmt.Scan(&size)
	for i := 0; i < size; i++ {
		fmt.Scan(&num)
		answer = (answer * num) % (1000000007)
	}

	fmt.Println(answer)
}
