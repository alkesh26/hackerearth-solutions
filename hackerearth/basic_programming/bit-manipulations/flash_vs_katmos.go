package main

import (
	"fmt"
	"sort"
)

type nums struct {
	id       int
	oneCount int
	num      int
}

type numbers []nums

func main() {
	var numberOfTests, size int
	fmt.Scanf("%v", &numberOfTests)

	for ; numberOfTests > 0; numberOfTests-- {
		fmt.Scan(&size)

		arr := make([]int, size)
		barray := make([]nums, size)

		for i := 0; i < size; i++ {
			fmt.Scan(&arr[i])
			barray[i].id = i + 1
			barray[i].num = arr[i]
			barray[i].oneCount = countOnes(arr[i])
		}

		sort.Sort(numbers(barray))

		for k := 0; k < size; k++ {
			fmt.Printf("%d ", barray[k].num)
		}

		fmt.Printf("\n")
	}

}

func countOnes(num int) int {
	ans := 0
	for num > 0 {
		ans += num & 1
		num >>= 1
	}

	return ans
}

func (n numbers) Len() int {
	return len(n)
}

func (n numbers) Less(i, j int) bool {
	if n[i].oneCount == n[j].oneCount {
		return n[i].id < n[j].id
	}
	return n[i].oneCount < n[j].oneCount
}

func (n numbers) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}
