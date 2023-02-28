import * as fs from "fs";

// util for input
// prettier-ignore
const lineit = (function* () { for (const line of fs.readFileSync(process.stdin.fd, "utf8").split("\n")) yield line.trim(); })();
// prettier-ignore
const wordit = (function* () { while (true) { let line = lineit.next(); if (line.done) break; for (const word of String(line.value).split(" ")) yield word; } })();
// prettier-ignore
const read = () => String((wordit.next()).value);

// main
const main = function () {
	// TODO edit this code, this code is for https://atcoder.jp/contests/practice/tasks/practice_1

	// param
	let nl: string[] = read().split("");
	let nl2: number[] = nl.map((v) => Number(v));

	// solve
	let count: number = nl2.reduce((a, b) => a + b, 0);
	// answer
	console.log(count);

	return;
};
main();
