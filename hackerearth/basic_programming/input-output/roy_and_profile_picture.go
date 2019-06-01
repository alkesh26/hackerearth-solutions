// Problem link:
// https://www.hackerearth.com/practice/basic-programming/input-output/basics-of-input-output/practice-problems/algorithm/roy-and-profile-picture/

package main

import (
	"fmt"
)

func main() {
	var minLength int
	fmt.Scanf("%d", &minLength)
	var numberOfTest int
	fmt.Scanf("%d", &numberOfTest)

	var length, width int
	for i := 0; i < numberOfTest; i++ {
		fmt.Scanf("%d %d", &length, &width)

		if length < minLength || width < minLength {
			fmt.Println("UPLOAD ANOTHER")
		} else if length == width {
			fmt.Println("ACCEPTED")
		} else {
			fmt.Println("CROP IT")
		}
	}
}
