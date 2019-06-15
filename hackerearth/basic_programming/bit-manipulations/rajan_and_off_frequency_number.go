package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	n = 2*n + 1

	arr := make([]int, n)

	odd := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
		odd ^= arr[i]
	}

	fmt.Println(odd)
}
