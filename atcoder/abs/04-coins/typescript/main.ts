// Copyright (c) 2023 Furukawa, Atsushi <atsushifx@aglabo.com>
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

/*
 * atcoder abs : coins
 */
import * as fs from "fs";

// stdin
const inputs: string = (fs.readFileSync(process.stdin.fd, "utf8"))
const inputArray: string[] = inputs.split(/\s/)
let inputIndex = 0;

// atcoder lib:stdin
function getNext(): string {
	return inputArray[inputIndex++];
}

function getNextNum(): number {
	return Number(getNext());
}

function getNums(N: number): number[] {
	const arr: number[] = [];
	for (let i: number = 0; i < N; ++i) {
		arr[i] = getNextNum();
	}
	return arr;
}

// lib:stdout
let outputBuffer: string = ""
function flush(): void {
	console.log(outputBuffer);
}

// exec main
solve();


// solve: solve problem main
function solve() {
	// Input
	const [a, b, c] = getNums(3);
	const x = getNextNum()

	// Answer
	let count: number = 0
	for (let i: number = 0; i <= a; ++i) {
		for (let j: number = 0; j <= b; ++j) {
			for (let k: number = 0; k <= c; ++k) {
				let sum: number = 500 * i + 100 * j + 50 * k
				if (sum == x) {
					count += 1
				}
				if (sum > x) {
					break
				}
			}
		}
	}

	// Out
	console.log(count)
}
