package main

import (
    "fmt"
    "sort"
)

//byLength is a type
type byLength []string

//We implement sort.Interface - Len, Less, and Swap on our type
// so we can use the sort package’s generic Sort function.
func (s byLength) Len() int {
    return len(s)
}
func (s byLength) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
    return len(s[i]) < len(s[j])
}

func main() {
    fruits := []string{"peach", "banana", "kiwi"}
    sort.Sort(byLength(fruits))
    fmt.Println(fruits)

//	ints := []int{7, 2, 4}
//    sort.Ints(ints)

}