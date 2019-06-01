// Problem link:
// https://www.hackerearth.com/practice/basic-programming/input-output/basics-of-input-output/practice-problems/algorithm/palindrome-check-2/

package main

import "fmt"

func main() {
	var str string
	fmt.Scanf("%s", &str)

	length := len(str)
	for i := 0; i < length/2; i++ {
		if str[i] != str[length-1-i] {
			fmt.Println("NO")
			return
		}
	}

	fmt.Println("YES")
}
