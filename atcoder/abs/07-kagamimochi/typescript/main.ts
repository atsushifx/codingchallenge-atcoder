// Copyright (c) 2023 Furukawa, Atsushi <atsushifx@aglabo.com>
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

// atcoder problem slolver
import * as fs from "fs"

/**
 * stdIO : parameter read and solve outputs
 *
 *
 */
class stdIO {
	// property
	// for stdin
	inputs: string = ""
	inputArray: string[]
	inputIndex: number = 0;

	// for stdout
	outputBuffer: string = "";

	// constructor]
	// readline from stdin and split this
	constructor() {
		this.inputs = (fs.readFileSync(process.stdin.fd, "utf8"))
		this.inputArray = this.inputs.split(/\s/)
		this.inputIndex = 0
	}

	// input meethod
	getNext(): string {
		return this.inputArray[this.inputIndex++];
	}

	getNextNum(): number {
		return Number(this.getNext());
	}

	getNums(N: number): number[] {
		const arr: number[] = [];
		for (let i: number = 0; i < N; ++i) {
			arr[i] = this.getNextNum();
		}
		return arr;
	}

	// bufferd output methods
	toString(...args: any[]): string {
		const buff: string[] = []
		args.forEach((arg: any) => {
			buff.push(this._toString(arg))
		})
		const buff2 = buff.join(' ')
		return buff2;
	}

	// 出力
	print(...args: any): void {
		this.outputBuffer += this.toString(...args)
	}

	p(...args: any[]): void {
		this.print(...args)
	}

	println(...args: any): void {
		this.print(...args);
		this.print("\n")
	}


	flush(): void {
		console.log(this.outputBuffer)
		this.outputBuffer = ""
	}

	// 引数を文字列か
	private _toString(arg1: any, quote: boolean = false): string {
		let buff = ""

		if (arg1 === null) {
			buff = "null"
		} else if (arg1 === "") {
		} else if (arg1 === "\n") {
			buff = "\n"
		} else if (Array.isArray(arg1)) {
			buff = this._arrayToString(arg1)
		} else {
			switch (typeof arg1) {
				case "undefined":
					buff = (typeof arg1);
					break
				case "string":
					if (quote) {
						buff = "'" + arg1 + "'";
					} else {
						buff = arg1
					}
					break
				default:
					buff = arg1.toString()
					break
			}
		}
		return buff;
	}

	// 配列を文字列に
	private _arrayToString(arr: any[]): string {
		const buff: string[] = []
		arr.map((e) => {
			buff.push(this._toString(e, true))
		})
		const buff2 = "[ " + buff.join(", ") + " ]"
		return buff2;
	}
}

// main routin
const io: stdIO = new stdIO()
solve()
io.flush()

// solver main()
function solve() {
	const N: number = io.getNextNum()
	const mochi: number[] = io.getNums(N)

	const kagamimochi: number[] = Array.from(new Set(mochi))

	io.p(kagamimochi.length)
}

