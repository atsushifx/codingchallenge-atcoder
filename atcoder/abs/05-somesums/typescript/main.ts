// Copyright (c) 2023 Furukawa, Atsushi <atsushifx@aglabo.com>
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

// atcoder solver : abs some sums

import * as fs from "fs";

/**
 * 標準入力から読み込み
 */
class sI {
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
}

/**
 * buffered Print
 */
class bP {
	static outputBuffer: string = "";

	// toString
	// methods
	static toString(...args: any[]): string {
		const buff: string[] = []
		args.forEach((arg: any) => {
			buff.push(this._toString(arg))
		})
		const buff2 = buff.join(' ')
		return buff2;
	}

	// 出力
	static print(...args: any): void {
		this.outputBuffer += this.toString(...args)
	}

	static p(...args: any[]): void {
		this.print(...args)
	}

	static println(...args: any): void {
		this.print(...args);
		this.print("\n")
	}


	static flush(): void {
		console.log(this.outputBuffer)
		this.outputBuffer = ""
	}

	// 引数を文字列か
	static _toString(arg1: any): string {
		let buff = ""
		if (arg1 === null) {
			buff = "null"
		} else if (arg1 == "") {
		} else if (arg1 == "\n") {
			buff = "\n"
		} else if (Array.isArray(arg1)) {
			buff = this._arrayToString(arg1)
		} else {
			switch (typeof arg1) {
				case "undefined":
					buff = (typeof arg1);
					break
				case "string":
					buff = "'" + arg1 + "'"
					break
				default:
					buff = arg1.toString()
					break
			}
		}
		return buff;
	}

	// 配列を文字列に
	static _arrayToString(arr: any[]): string {
		const buff: string[] = []
		arr.map((e) => {
			buff.push(this._toString(e))
		})
		const buff2 = "[ " + buff.join(", ") + " ]"
		return buff2;
	}
}

//

// exec main
solve();
flush()



// solver main function
function solve(): void {
	const [N, A, B]: number[] = getNums(3);

	console.log(N, [A, B]);
}
