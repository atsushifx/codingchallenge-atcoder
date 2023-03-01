// Copyright (c) 2023 Furukawa, Atsushi <atsushifx@aglabo.com>
// 
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
var twoSum = function (nums, target) {
	dic = []
	for (let i = 0; i < nums.length - 1; i++) {
		const n = nums[i]
		let i1 = dic[target - n]
		if (i1 !== undefined) {
			return [i1, i];
		}
		dic[n] = i
	}
	console.log(dic)
	return [];
};

console.log(twoSum([2, 7, 11, 8], 9))
