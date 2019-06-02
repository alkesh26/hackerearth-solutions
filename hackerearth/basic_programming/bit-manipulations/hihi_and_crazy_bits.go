// Problem link:
// https://www.hackerearth.com/practice/basic-programming/bit-manipulation/basics-of-bit-manipulation/practice-problems/algorithm/hihi-and-crazy-bits/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	var numberOfTest, n int
	fmt.Fscan(reader, &numberOfTest)

	for numberOfTest > 0 {
		fmt.Fscan(reader, &n)

		if n%2 == 0 {
			fmt.Fprintf(writer, "%d\n", n+1)
		} else {
			fmt.Fprintf(writer, "%d\n", n|1<<checkToggle(n))
		}
		numberOfTest--
	}

	writer.Flush()
}

func checkToggle(n int) uint {
	var i uint
	for {
		if (n & (1 << i)) == 0 {
			return i
		}
		i++
	}
}
