package main

import (
	"fmt"
)

func average(a []int) float64{
	if len(a) == 0 {
		return 0
	}

	var sum int 

	for i := 0; i < len(a); i++ {
		sum += a[i] 
	}

	return float64(sum) / float64(len(a))
}

func main(){
	fmt.Print("Average value of array is ")
	fmt.Println(average([]int{1, 2, 3, 4, 5, 6}))
}
