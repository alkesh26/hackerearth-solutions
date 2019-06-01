// Problem link:
// https://www.hackerearth.com/practice/basic-programming/input-output/basics-of-input-output/practice-problems/algorithm/count-divisors/

package main

import (
	"fmt"
)

func main() {
	var l, r, k int
	fmt.Scan(&l)
	fmt.Scan(&r)
	fmt.Scan(&k)

	if l == r && l%k != 0 {
		fmt.Println(0)
		return
	} else if l == r && l%k == 0 {
		fmt.Println(1)
		return
	}

	if l%k != 0 {
		l = k - (l % k) + l
	}
	fmt.Println((r-l)/k + 1)
}
