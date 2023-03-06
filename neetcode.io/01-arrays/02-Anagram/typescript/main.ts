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
	let s: string
	let t: string

	s = "anagram";
	t = "nagaram"
	ans = solver.isAnagram(s, t)
	console.log(s + ", " + t + " answer : " + ans)

	s = "rat";
	t = "car"
	ans = solver.isAnagram(s, t)
	console.log(s + ", " + t + " answer : " + ans)
}

