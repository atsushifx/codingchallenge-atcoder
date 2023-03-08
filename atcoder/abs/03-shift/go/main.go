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
	var size int = getNextInt()
	var nums []int = getNums(size)

	var isEven bool = true
	var count int = 0
	for isEven {
		for i := range nums {
			if (nums[i] % 2) != 0 {
				isEven = false
				break
			}
			nums[i] /= 2
		}
		if isEven {
			count++
		}
	}

	fmt.Fprintf(Stdout, "%d\n", count)
}
