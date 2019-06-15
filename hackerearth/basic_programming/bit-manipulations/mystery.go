package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	count := 0
	check := false
	for {
		num := -1
		fmt.Fscan(reader, &num)

		if num == -1 {
			break
		}

		if count > 1 && num == 0 {
			check = true
			continue
		}

		if count > 49 && check {
			break
		}

		fmt.Fprintf(writer, "%d\n", (num & (-num)))
		count++

		writer.Flush()
	}
}
