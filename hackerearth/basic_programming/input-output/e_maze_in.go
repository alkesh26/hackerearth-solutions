// Problem link:
// https://www.hackerearth.com/practice/basic-programming/input-output/basics-of-input-output/practice-problems/algorithm/e-maze-in-1aa4e2ac/

package main

import "fmt"

func main() {
	var command string
	fmt.Scan(&command)

	a, b := 0, 0
	for i := 0; i < len(command); i++ {
		switch command[i] {
		case 'L':
			a--
		case 'R':
			a++
		case 'D':
			b--
		case 'U':
			b++
		}
	}

	fmt.Println(a, b)
}
