// Copyright (c) 2023 Furukawa, Atsushi <atsushifx@aglabo.com>
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

// modules
import (
	"fmt"
)

// main
func main() {
	var ans bool

	var arr1 []int = []int{1, 2, 3, 1}
	ans = containsDuplicate(arr1)
	fmt.Printf("%v %v\n", arr1, ans)

	var arr2 []int = []int{1, 2, 3, 8}
	ans = containsDuplicate(arr2)
	fmt.Printf("%v %v\n", arr2, ans)
}
