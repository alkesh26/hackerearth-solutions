package main

import "fmt"

func main() {
	var n, num int
	fmt.Scan(&n)

	for t := 0; t < n; t++ {
		fmt.Scan(&num)

		count := 0

		if num&(num-1) == 0 {
			count = 1
		} else {
			for num > 0 {
				count++
				num = num & (num - 1)
			}
		}

		fmt.Println(count)
	}
}
