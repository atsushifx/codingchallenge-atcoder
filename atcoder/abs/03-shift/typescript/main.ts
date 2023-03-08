// Copyright (c) 2023 Furukawa, Atsushi <atsushifx@aglabo.com>
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

/*
 * atcoder typescript template
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
function print(out: string | number | bigint): void;
function print<T>(out: Array<T>, separator: string): void;
function print<T>(out: string | number | bigint | Array<T>, separator?: string) {
	if (Array.isArray(out)) {
		outputBuffer += out.join(separator);
	} else {
		outputBuffer += out;
	}
}

function println(out: string | number | bigint): void;
function println<T>(out: Array<T>, separator: string): void;
function println<T>(out: string | number | bigint | Array<T>, separator?: string) {
	if (Array.isArray(out)) {
		print(out, separator || "");
	} else {
		print(out);
	}
	print("\n");
}

function flush(): void {
	console.log(outputBuffer);
}

// main
solve();

// main
function solve() {
	// Input
	const N: number = getNextNum();
	let nums: number[] = getNums(N)

	let isEven: boolean = true
	let count: number = 0
	while (isEven) {
		nums = nums.map(
			(num): number => {
				if ((num % 2) != 0) {
					isEven = false
				}
				return num / 2
			}
		)
		if (isEven) {
			count++;
		}
	}

	// Out
	console.log(`${count}`);
}
