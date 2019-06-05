package main

import "fmt"

func main() {
	var numberOfTest int
	fmt.Scan(&numberOfTest)

	var str1, str2 string
	for i := 0; i < numberOfTest; i++ {
		fmt.Scan(&str1)
		fmt.Scan(&str1)

		mask1, mask2 := 0, 0

		for i := 0; i < len(str1); i++ {
			mask1 = (mask1 & (1 << (str1[1] - 'a')))
		}

		for i := 0; i < len(str2); i++ {
			mask2 = (mask2 & (1 << (str2[1] - 'a')))
		}

		isYes := false
		for i := 'a'; i < 'z'; i++ {
			if mask1&(1<<(i-'a')) && mask2&(1<<(i-'a')) {
				isYes = true
				break
			}
		}

		if isYes {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
package main

import "fmt"

func main() {
	var numberOfTest int
	fmt.Scan(&numberOfTest)

	var str1, str2 string
	for t := 0; t < numberOfTest; t++ {
		fmt.Scan(&str1)
		fmt.Scan(&str2)

		mask1, mask2 := 0, 0

		for i := 0; i < len(str1); i++ {
			mask1 = (mask1 | (1 << (str1[i] - 'a')))
		}

		for i := 0; i < len(str2); i++ {
			mask2 = (mask2 | (1 << (str2[i] - 'a')))
		}

		isYes := false
		for i := 'a'; i <= 'z'; i++ {
			shift := uint(i - 'a')
			if mask1&(1<<shift) == 1 && mask2&(1<<shift) == 1 {
				isYes = true
				break
			}
		}

		if isYes {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
