package main

import (
 "fmt"
 "os"
 "sort"
 "strings"
 "time"
)

type Person struct {
 firstName string
 lastName  string
 birthDay  time.Time
}

type People []Person

func (s People) Len() int {
 return len(s)
}

func (s People) Less(i, j int) bool {
 if s[i].birthDay.Equal(s[j].birthDay) {
	 if strings.ToLower(s[i].firstName) == strings.ToLower(s[j].firstName) {
		 return strings.ToLower(s[i].lastName) < strings.ToLower(s[j].lastName)
	 }
	 return strings.ToLower(s[i].firstName) < strings.ToLower(s[j].lastName)
 }
 return s[i].birthDay.After(s[j].birthDay)
}

func (s People) Swap(i, j int) {
 s[i], s[j] = s[j], s[i]
}

func (s People) String() string {
 var result string
 for _, i := range s {
	 result += fmt.Sprintf("%s %s %s\n", i.firstName, i.lastName, i.birthDay)
 }
 return result
}

func main() {
 ivanIvanovDate1, err := time.Parse("2006-Jan-02", "2005-Aug-10")
 artiomIvanovDate, err2 := time.Parse("2006-Jan-02", "2005-Aug-10")
 ivanIvanovDate2, err3 := time.Parse("2006-Jan-02", "2003-Aug-05")
 if err != nil || err2 != nil || err3 != nil {
	 fmt.Println(err)
	 os.Exit(1)
 }

 p := People{
	 {"Ivan", "Ivanov", ivanIvanovDate1},
	 {"Artiom", "Ivanov", artiomIvanovDate},
	 {"Ivan", "Ivanov", ivanIvanovDate2},
 }

 sort.Sort(p)
 fmt.Println(p)
}
