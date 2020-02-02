package main

import (
	// "./custompkg"
	"github.com/mohit-kumar-sharma/go-template-documentation/custompkg"
	"fmt"
)

func main() {
	s := []int{1,2,3,4,5}
	avg := custompkg.CustomAvg(s)
	fmt.Println(avg)
}
