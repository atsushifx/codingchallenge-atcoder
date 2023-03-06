// Copyright (c) 2023 Furukawa, Atsushi <atsushifx@aglabo.com>
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

// modules
import (
	"reflect"
)

// problem solver
func isAnagram(s string, t string) bool {
	var sH = strHist(s)
	var tH = strHist(t)
	return reflect.DeepEqual(sH, tH)
}

func strHist(str string) []int {
	var h = make([]int, 128)

	for _, c := range str {
		h[c]++
	}
	return h
}
