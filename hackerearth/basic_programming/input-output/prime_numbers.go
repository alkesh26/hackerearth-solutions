// Problem link:
// https://www.hackerearth.com/practice/basic-programming/input-output/basics-of-input-output/practice-problems/algorithm/prime-number-8/

package main

import "fmt"

func main() {
	var num int
	fmt.Scanf("%d", &num)

	if num == 2 {
		fmt.Println(2)
	} else {
		isPrimes := make([]bool, num)
		isPrimes[0], isPrimes[1] = true, true
		for i := 2; i < num; i++ {
			if !isPrimes[i] {
				fmt.Printf("%d ", i)
			}
			var prod int
			for j := 2; prod < num; j++ {
				prod = i * j
				if prod < num {
					isPrimes[prod] = true
				}
			}
		}
	}
}
