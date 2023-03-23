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
const N = 10 ^ 5

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

func getNextTPoint() TPoint {
	var tp = new(TPoint)

	tp.t = getNextNum()
	tp.x = getNextNum()
	tp.y = getNextNum()
	return *tp
}

// main routin
// global var
var TP0 (TPoint) = TPoint{0, 0, 0}

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

func abs64(x int64) int64 {
	if x >= 0 {
		return x
	} else {
		return -x
	}
}

func canMoveList(tp *[N + 1]TPoint, n int64) bool {
	for i := (int64)(1); i <= n; i++ {
		var dt int64 = abs64(tp[i].t-tp[i-1].t) -
			(abs64(tp[i].x-tp[i-1].x) + abs64(tp[i].y-tp[i-1].y))

		if dt < 0 {
			return false
		}
		if (dt % 2) > 0 {
			return false
		}

	}
	return true
}

// problem solver
func solve() {
	// initialize
	var tplist [N + 1]TPoint
	tplist[0] = TP0
	// input
	var n int64 = getNextNum()
	for i := (int64)(1); i <= n; i++ {
		tplist[i] = getNextTPoint()
	}

	// solve
	var canMove bool = canMoveList(&tplist, n)

	// output
	if canMove {
		fmt.Fprintf(Stdout, "Yes\n")
	} else {
		fmt.Fprintf(Stdout, "No\n")
	}
}
