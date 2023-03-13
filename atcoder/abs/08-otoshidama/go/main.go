// Copyright (c) 2023 Furukawa, Atsushi <atsushifx@aglabo.com>
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

// atcoder problem solver
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Global Variables
const IOBUFFSIZE = 10 ^ 6

var Stdin = bufio.NewScanner(os.Stdin)
var Stdout = bufio.NewWriter(os.Stdout)

// atcoder lib
func getNext() string {
	Stdin.Scan()
	return Stdin.Text()
}

func getNextInt() int64 {
	var i int64
	var e error
	i, e = strconv.ParseInt(getNext(), 10, 64)
	if e != nil {
		panic(e)
	}
	return i
}

// initialize
func initialize() {
	stdinBuffer := make([]byte, IOBUFFSIZE)
	Stdin.Buffer(stdinBuffer, IOBUFFSIZE)
	Stdin.Split(bufio.ScanWords)
}

// main
func main() {
	initialize()
	defer Stdout.Flush()
	solve()
	Stdout.Flush()
}

// problem solver
func solve() {
	// input
	var N int64 = getNextInt()
	var Y int64 = getNextInt()

	// fmt.Fprintln(Stdout, N, Y)

	// solve
	var isFound bool = false
	var i int64 = 0
	var j int64 = 0
	var k int64 = 0
	for i = 0; i <= N; i++ {
		for j = 0; j <= N-i; j++ {
			k = N - (i + j)
			var sum int64 = i*10000 + j*5000 + k*1000

			// fmt.Fprintf(Stdout, "[%d,%d,%d]=%d, %v\n", i, j, k, sum, sum > Y)
			if sum == Y {
				isFound = true
				break
			}
			if sum > Y {
				break
			}
		}
		if isFound {
			break
		}
	}

	// output
	if isFound {
		fmt.Fprintf(Stdout, "%d %d %d\n", i, j, k)
	} else {
		fmt.Fprintln(Stdout, "-1 -1 -1")
	}
}
