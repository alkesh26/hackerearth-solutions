package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	var tests int
	fmt.Fscan(reader, &tests)
	for tests > 0 {
		var n int
		fmt.Fscan(reader, &n)
		array := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &array[i])
		}
		givenSum := 0
		fmt.Fscan(reader, &givenSum)
		if givenSum == 0 {
			fmt.Fprintf(writer, "%v\n", "YES")
		} else {
			found := false
			for i := 1; i < (1 << uint(n)); i++ {
				sum := 0
				for j := 0; j < N; j++ {
					if (i & (1 << uint(j))) != 0 {
						sum += array[j]
					}
				}
				if sum == givenSum {
					fmt.Fprintf(writer, "%v\n", "YES")
					found = true
					break

				}

			}
			if !found {
				fmt.Fprintf(writer, "%v\n", "NO")
			}
		}
		writer.Flush()
		tests--
	}
}
