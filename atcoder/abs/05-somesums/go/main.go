// Copyright (c) 2023 Furukawa, Atsushi <atsushifx@aglabo.com>
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

// atcoder problem solver
// problem; abs / some sums
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

// solver's utilty
func addNums(num int) int {
	var sum int = 0
	for num > 0 {
		sum += num % 10
		num /= 10
	}
	return sum
}

// problem solver
func solve() {
	var N, A, B int
	N = getNextInt()
	A = getNextInt()
	B = getNextInt()

	var total int = 0
	for i := 1; i <= N; i++ {
		var sum int = addNums(i)
		if A <= sum && sum <= B {
			total += i
		}
	}
	fmt.Fprintln(Stdout, total)
}
