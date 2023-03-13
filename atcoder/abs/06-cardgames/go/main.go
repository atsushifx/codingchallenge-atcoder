// Copyright (c) 2023 Furukawa, Atsushi <atsushifx@aglabo.com>
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	var N int = getNextInt()
	var cards []int = getNums(N)

	sort.Slice(cards, func(i, j int) bool { return cards[i] > cards[j] })

	// calc points
	var pt int = 0
	for i := 0; i < N; i++ {
		if (i % 2) == 0 {
			pt += cards[i]
		} else {
			pt -= cards[i]
		}
	}
	fmt.Fprintln(Stdout, pt)
}
