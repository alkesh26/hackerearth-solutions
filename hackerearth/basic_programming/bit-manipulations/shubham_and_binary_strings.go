package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var numberOfTest int
	fmt.Fscan(reader, &numberOfTest)

	for ; numberOfTest > 0; numberOfTest-- {
		var n int
		var str string
		fmt.Fscan(reader, &n)
		fmt.Fscan(reader, &str)

		count := 0
		for i := 0; i < n; i++ {
			if str[i] == '0' {
				count++
			}
		}

		if count != n {
			fmt.Fprintf(writer, "%d\n", count)
		} else {
			fmt.Fprintf(writer, "%d\n", 1)
		}
	}

	writer.Flush()
}
