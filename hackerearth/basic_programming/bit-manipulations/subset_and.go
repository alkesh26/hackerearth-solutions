package main

import (
	"fmt"
)

func main() {
	var numberOfTests, n, z int
	fmt.Scanf("%v", &numberOfTests)

	for ; numberOfTests > 0; numberOfTests-- {
		fmt.Scanf("%v%v\n", &z, &n)

		found := false
		allAnd := z
		for i := 0; i < n; i++ {
			a := 0
			if i == n-1 {
				fmt.Scanf("%v\n", &a)
			} else {
				fmt.Scanf("%v", &a)
			}

			if (z & a) == 0 {
				found = true
			}
			allAnd = allAnd & a

		}
		if allAnd == 0 || found {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}

	}

}
