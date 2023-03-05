// Copyright (c) 2023 Furukawa, Atsushi <atsushifx@aglabo.com>
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

// import solver
import solver from './solver.js';

main()
// main
function main() {
	let ans: boolean;

	const arr1: number[] = [1, 2, 3, 1]
	ans = solver.containsDuplicate(arr1)
	console.log(arr1, " answer : " + ans)

	const arr2: number[] = [1, 2, 3, 4]
	ans = solver.containsDuplicate(arr2)
	console.log(arr2, " answer : " + ans)
}

