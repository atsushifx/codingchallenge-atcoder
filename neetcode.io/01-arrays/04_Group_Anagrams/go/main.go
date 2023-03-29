// Copyright (c) 2023 Furukawa, Atsushi <atsushifx@aglabo.com>
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT
package main

import (
	"fmt"
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	var hash = make(map[string][]string, len(strs))
	for _, s := range strs {
		var h = strHash(s)
		hash[h] = append(hash[h], s)
	}
	var ret [][]string = values((hash))
	return ret
}

func values(m map[string][]string) [][]string {
	vs := [][]string{}
	for _, v := range m {
		vs = append(vs, v)
	}
	return vs
}

// get sorted string
func strHash(str string) string {
	var chars []string = strings.Split(str, "")
	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})
	var hash = strings.Join(chars, "")
	return hash
}

// main roution
func main() {
	input := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	output := groupAnagrams(input)

	fmt.Println(output)
}
