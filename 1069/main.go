/*
Track the current repetition length while scanning the DNA sequence once.

For each character, either extend the current repetition or reset the counter.
The maximum repetition length is updated continuously during iteration.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var s string
	fmt.Fscan(reader, &s)

	longestSize := 1
	currentSize := 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			currentSize++
		} else {
			currentSize = 1
		}
		longestSize = max(longestSize, currentSize)
	}

	fmt.Fprintln(writer, longestSize)
}
