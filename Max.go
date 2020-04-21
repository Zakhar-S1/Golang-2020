package main 

import (
	"fmt"
	"unicode/utf8"
)

func max(a [] string) string{

	if len(a) == 0 {
		return "Sorry, but you have empty slice."
	}

	var s string = a[0]
	
	for i := 1; i < len(a); i++{
		if utf8.RuneCountInString(a[i]) > utf8.RuneCountInString(s){
			s = a[i]
		}
	}
	return s
}

func main(){
	fmt.Print("The longest word from the slice of strings - ")
	fmt.Println(max([] string{"one","two","three"}))
	fmt.Print("the longest word from the slice of strings, if there are more than one - ")
	fmt.Println(max([] string{"one","two"}))
}
