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
	// initialize
	var Keywords = []string{
		"dream",
		"dreamer",
		"erase",
		"eraser",
	}

	// input
	var S string = getNext()

	// solve
	var match bool = isMatch(S, Keywords)

	// output
	if match {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}

// solver function
func isMatch(str string, keywords []string) bool {
	var ret bool = false
	for _, key := range keywords {
		if str == key {
			ret = true
			break
		} else if strings.HasSuffix(str, key) {
			var str2 string = fmt.Sprint(str[:len(str)-len(key)])
			ret = isMatch(str2, keywords)
			break
		} else {

		}
	}
	return ret
}
