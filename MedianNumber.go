package main

import (
	"fmt"
	"sort"
)

func medianNumber(i []int) float64 {
	
	sort.Ints(i)

	if len(i)%2 == 0 {
		return float64(i[len(i)/2] + i[len(i)/2-1]) / 2
	}

	return float64(i[len(i)/2])
}

func main() {

	fmt.Println(medianNumber([]int{1, 2, 3, 4, 5}))
	fmt.Println(medianNumber([]int{3, 3, 5, 6, 9, 11}))
	fmt.Println(medianNumber([]int{1, 2, 2, 6}))
	fmt.Println(medianNumber([]int{4, 3, 8, 5}))
	fmt.Println(medianNumber([]int{9, 5, 1}))

}
