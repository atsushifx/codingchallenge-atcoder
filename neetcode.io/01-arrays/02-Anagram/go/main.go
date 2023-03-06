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
	var s, t string

	s = "anagram"
	t = "nagaram"
	ans = isAnagram(s, t)
	fmt.Printf("[%s , %s] %v\n", s, t, ans)

	s = "rat"
	t = "car"
	ans = isAnagram(s, t)
	fmt.Printf("[%s , %s] %v\n", s, t, ans)
}
