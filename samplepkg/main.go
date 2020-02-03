// Package sample is used to show a sample
package sample

import "fmt"

// Sayhi is used to say hi to a person
func SayHi(name string) string {
	return fmt.Sprintf("Hi %v", name)
}
