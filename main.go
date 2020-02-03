package main

import (
	// "./custompkg"
	"fmt"

	"go/build"

	"os"

	"github.com/mohit-kumar-sharma/go-template-documentation/custompkg"

	sample "github.com/mohit-kumar-sharma/go-template-documentation/samplepkg"
)

func main() {
	s := []int{1, 2, 3, 4, 5}
	avg := custompkg.CustomAvg(s)
	fmt.Println(avg)
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}
	fmt.Println(gopath)

	fmt.Println(sample.SayHi("Mohit"))

}
