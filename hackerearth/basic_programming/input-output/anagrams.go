// Problem link:
// https://www.hackerearth.com/practice/basic-programming/input-output/basics-of-input-output/practice-problems/algorithm/anagrams-651/

package main

import (
	"fmt"
)

func main() {
	var numberOfTests int
	fmt.Scanf("%d", &numberOfTests)

	for i := 0; i < numberOfTests; i++ {
		var a, b string
		fmt.Scanf("%s", &a)
		fmt.Scanf("%s", &b)

		characterArray := make([]int, 26)

		for _, c := range a {
			characterArray[int(c)-'a']++
		}

		for _, c := range b {
			characterArray[int(c)-'a']--
		}

		numberOfDeletions := 0
		for i := 0; i < 26; i++ {
			if characterArray[i] < 0 {
				numberOfDeletions += characterArray[i] * -1
			} else {
				numberOfDeletions += characterArray[i]
			}
		}
		fmt.Println(numberOfDeletions)
	}
}
