package main

import (
	"fmt"
	"sort"
)

func printSorted(a map[int]string){
	key := make([]int, 0, len(a))
	
	for i := range a {
		key = append(key, i)
	}
	
	sort.Ints(key)
	
	for _, i := range key {
		fmt.Print(a[i], " ")
	}

	fmt.Println()
}

func main(){
	fmt.Print("The first sorted map - ")
	printSorted(map [int]string{2: "a", 0: "b", 1: "c"})
	
	fmt.Print("The second sorted map - ")
	printSorted(map [int]string{10: "aa", 0: "bb", 500: "cc"})
}
