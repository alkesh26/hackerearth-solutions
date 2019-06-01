// Problem link:
// https://www.hackerearth.com/practice/basic-programming/input-output/basics-of-input-output/practice-problems/algorithm/modify-the-string/

package main

import (
	"fmt"
)

func main() {
	var str string
	fmt.Scan(&str)
	bytes := make([]byte, len(str))
	for i := 0; i < len(str); i++ {
		if str[i] >= 'A' && str[i] <= 'Z' {
			bytes[i] = str[i] + 'a' - 'A'
		} else if str[i] >= 'a' && str[i] <= 'z' {
			bytes[i] = str[i] + 'A' - 'a'
		}
	}

	fmt.Println(string(bytes))
}
