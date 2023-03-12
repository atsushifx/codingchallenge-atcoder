// Copyright (c) 2023 Furukawa, Atsushi <atsushifx@aglabo.com>
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

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

func getNextInt() int {
	var i int
	var e error
	i, e = strconv.Atoi(getNext())
	if e != nil {
		panic(e)
	}
	return i
}

func getNums(size int) []int {
	var buff []int = make([]int, size)
	for i := range buff {
		buff[i] = getNextInt()
	}
	return buff
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
	var a int = getNextInt()
	var b int = getNextInt()
	var c int = getNextInt()

	var x = getNextInt()

	// solve
	var count int = 0

	for i := 0; i <= a; i++ {
		for j := 0; j <= b; j++ {
			for k := 0; k <= c; k++ {
				var sum = 500*i + 100*j + 50*k
				if sum == x {
					count++
					// fmt.Fprintf(Stdout, "[%d %d %d]\n", i, j, k)
				}
			}
		}
	}

	// Out
	fmt.Fprintln(Stdout, count)

}
