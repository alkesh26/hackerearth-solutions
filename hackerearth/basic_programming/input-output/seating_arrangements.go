// Problem link:
// https://www.hackerearth.com/practice/basic-programming/input-output/basics-of-input-output/practice-problems/algorithm/seating-arrangement-1/

package main

import "fmt"

func main() {
	var numberOfTests int
	fmt.Scanf("%d", &numberOfTests)

	var num int
	for i := 0; i < numberOfTests; i++ {
		fmt.Scanf("%d", &num)
		rem := num % 12
		switch rem {
		case 0:
			fmt.Println(num-11, "WS")
		case 1:
			fmt.Println(num+11, "WS")
		case 2:
			fmt.Println(num+9, "MS")
		case 3:
			fmt.Println(num+7, "AS")
		case 4:
			fmt.Println(num+5, "AS")
		case 5:
			fmt.Println(num+3, "MS")
		case 6:
			fmt.Println(num+1, "WS")
		case 7:
			fmt.Println(num-1, "WS")
		case 8:
			fmt.Println(num-3, "MS")
		case 9:
			fmt.Println(num-5, "AS")
		case 10:
			fmt.Println(num-7, "AS")
		case 11:
			fmt.Println(num-9, "MS")
		}
	}
}
