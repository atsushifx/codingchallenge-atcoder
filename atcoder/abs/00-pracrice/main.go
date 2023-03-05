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
	var a, b, c int
	a = getNextInt()
	b = getNextInt()
	c = getNextInt()
	var sum = a + b + c

	var inputs string = getNext()
	fmt.Fprintf(Stdout, "%d %s\n", sum, inputs)
}
