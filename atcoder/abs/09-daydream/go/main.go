// Copyright (c) 2023 Furukawa, Atsushi <atsushifx@aglabo.com>
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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
	var str string = getNext()

	// solve
	var matcher string = `^(dream|dreamer|erase|eraser)*$`
	var match bool = regexp.MustCompile(matcher).Match([]byte(str))

	// output
	if match {
		fmt.Fprintln(Stdout, "YES")
	} else {
		fmt.Fprintln(Stdout, "NO")
	}
}
