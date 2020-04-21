package main 

import (
	"fmt"
)

func reverse(a []int64) []int64{
	slice_1 := make([]int64, 0, len(a))
	for i := len(a) - 1; i >= 0; i-- {
		slice_1 = append(slice_1, a[i])
	}

	return slice_1
}

func main(){
	var slice_2 = [] int64{1, 2, 5, 15}
	fmt.Print("The reverse versing of slice: ")
	fmt.Println(reverse(slice_2))
}
