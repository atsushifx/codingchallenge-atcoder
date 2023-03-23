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
	"strings"
)

// Global Variables
const IOBUFFSIZE = 10 ^ 6

var Stdin = bufio.NewScanner(os.Stdin)
var Stdout = bufio.NewWriter(os.Stdout)

// atcoder lib
func getNext() string {
	Stdin.Scan()
	return strings.TrimSpace(Stdin.Text())
}

func getNextNum() int64 {
	var i int64
	var e error
	i, e = strconv.ParseInt(getNext(), 10, 64)
	if e != nil {
		panic(e)
	}
	return i
}

// solver
type TPoint struct {
	t int64
	x int64
	y int64
}

func getNextTPoint() *TPoint {
	var tp = new(TPoint)

	tp.t = getNextNum()
	tp.x = getNextNum()
	tp.y = getNextNum()
	return tp
}

// main routin
// global var
var TP0 (*TPoint) = &TPoint{0, 0, 0}

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
	var n int64 = getNextNum()
	var tplist = [](*TPoint){TP0}

	for i := (int64)(1); i <= n; i++ {

	}

	fmt.Fprintf(Stdout, "%v\n", tplist)
}
