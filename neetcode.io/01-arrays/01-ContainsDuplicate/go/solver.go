// Copyright (c) 2023 Furukawa, Atsushi <atsushifx@aglabo.com>
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

// modules

// problem solver
func containsDuplicate(nums []int) bool {
	var nMap = map[int]int{}
	for i := 0; i < len(nums); i++ {
		nMap[nums[i]]++
		if nMap[nums[i]] >= 2 {
			return true
		}
	}
	return false
}
