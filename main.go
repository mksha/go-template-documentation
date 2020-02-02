package main

import (
	"./custompkg"
	"fmt"
)

func main() {
	s := []int{1,2,3,4,5}
	avg := custompkg.CustomAvg(s)
	fmt.Println(avg)
}
