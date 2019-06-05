package main

import "fmt"

func main() {
	var numberOfTest, n, x int
	var max, maxIndex int
	fmt.Scan(&numberOfTest)

	for i := 0; i < numberOfTest; i++ {
		fmt.Scan(&n)
		arr := make([]int, 32)

		for i := 0; i < n; i++ {
			fmt.Scan(&x)
			j := 0
			for x > 0 {
				arr[j] += x % 2
				x /= 2
				j++
			}
		}
		max = 0
		maxIndex = 0
		for i := 0; i < 32; i++ {
			if arr[i] > max {
				max = arr[i]
				maxIndex = i
			}
		}

		fmt.Println(maxIndex)
	}
}
