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
	var N int = getNextInt()
	var mochi []int = getNums(N)

	// solve
	var kagamimochi []int = arrayUniq(mochi)

	// output
	// fmt.Fprintf(Stdout, "%d\n%v\n%v\n", N, mochi, kagamimochi)
	fmt.Fprintln(Stdout, len(kagamimochi))
}

// uniq
func arrayUniq(src []int) (dist []int) {
	m := map[int]bool{}

	for _, v := range src {
		if !m[v] {
			m[v] = true
			dist = append(dist, v)
		}
	}
	return dist
}
