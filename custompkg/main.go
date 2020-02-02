// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package custompkg is an documentation example.
package custompkg

// use customsum function to get the sum of a slice elements.
func customsum(s []int) int {
	var sum int
	for _, v := range s {
		sum += v
	}

	return sum
}

// CustomAvg used to get the average of
// elements of a slice
func CustomAvg(s []int) float64 {
	var avg float64

	sum := customsum(s)
	n := len(s)

	avg = float64(sum / n)

	return avg
}
