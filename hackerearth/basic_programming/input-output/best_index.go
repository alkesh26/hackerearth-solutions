// Best Index Hackerearth 
// Problem link:
// https://www.hackerearth.com/practice/basic-programming/input-output/basics-of-input-output/practice-problems/algorithm/best-index-1-45a2f8ff/

package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	arr := make([]int, n)
	prefixSum := make([]int, n+1)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	for i := 1; i <= n; i++ {
		prefixSum[i] = arr[i-1] + prefixSum[i-1]
	}

	max := math.MinInt32
	for i := 0; i < n; i++ {
		k := i

		for j := 2; j+k < n; j++ {
			k = j + k
		}

		res := prefixSum[k+1] - prefixSum[i]

		if res > max {
			max = res
		}
	}

	fmt.Println(max)
}
